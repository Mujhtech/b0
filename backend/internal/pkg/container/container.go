package container

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

type Container struct {
	client *client.Client
}

func New() (*Container, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	return &Container{client: cli}, nil
}

func (c *Container) Close() error {
	return c.client.Close()
}

func (c *Container) GetClient() *client.Client {
	return c.client
}

func (c *Container) PullImage(ctx context.Context, imageRef string) error {
	reader, err := c.client.ImagePull(ctx, imageRef, image.PullOptions{})
	if err != nil {
		return err
	}

	defer reader.Close()

	buf := new(strings.Builder)

	_, err = io.Copy(buf, reader)

	if err != nil {
		return err
	}

	return nil
}

func (c *Container) GetProjectContainer(ctx context.Context, id string) (*types.ContainerJSON, error) {
	container, _, err := c.client.ContainerInspectWithRaw(ctx, id, true)
	return &container, err
}

func (c *Container) IsContainerExist(ctx context.Context, opts FilterContainer) (bool, error) {

	filter := filters.NewArgs()

	if opts.Name != "" {
		filter.Add("name", opts.Name)
	}

	if opts.ID != "" {
		filter.Add("id", opts.ID)
	}

	if opts.Port != "" {
		filter.Add("port", opts.Port)
	}

	containers, err := c.client.ContainerList(ctx, container.ListOptions{
		Filters: filter,
		All:     true,
	})
	if err != nil {
		return false, err
	}

	return len(containers) > 0, nil
}

func (c *Container) GetContainerList(ctx context.Context, opts FilterContainer) ([]types.Container, error) {
	filter := filters.NewArgs()

	if opts.Name != "" {
		filter.Add("name", opts.Name)
	}

	if opts.ID != "" {
		filter.Add("id", opts.ID)
	}

	if opts.Port != "" {
		filter.Add("port", opts.Port)
	}

	containers, err := c.client.ContainerList(ctx, container.ListOptions{
		Filters: filter,
		All:     true,
	})
	if err != nil {
		return nil, err
	}

	return containers, nil
}

func (c *Container) CreateContainer(ctx context.Context, opts CreateContainer) (string, error) {

	commands := []string{"/bin/sh", "-c"}

	if len(opts.Command) > 0 {
		commands = append(commands, opts.Command...)
	}

	resp, err := c.client.ContainerCreate(ctx, &container.Config{
		OpenStdin:    true,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
		WorkingDir:   "/app",
		Image:        opts.Image,
		Cmd:          commands,
		ExposedPorts: nat.PortSet{
			nat.Port(fmt.Sprintf("%s/tcp", opts.Port)): struct{}{},
		},
		Labels: opts.Labels,
	}, &container.HostConfig{
		Binds: []string{
			"/var/run/docker.sock:/var/run/docker.sock",
		},
		PortBindings: map[nat.Port][]nat.PortBinding{
			nat.Port(fmt.Sprintf("%s/tcp", opts.Port)): {
				{
					// HostIP:   "0.0.0.0",
					HostPort: opts.Port,
				},
			},
		},
		Resources: container.Resources{
			CPUQuota:  100000,
			CPUPeriod: 100000,
		},
	}, &network.NetworkingConfig{}, nil, opts.Name)
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}

func (c *Container) StartContainer(ctx context.Context, id string) error {
	return c.client.ContainerStart(ctx, id, container.StartOptions{})
}

func (c *Container) StopContainer(ctx context.Context, id string) error {
	return c.client.ContainerStop(ctx, id, container.StopOptions{})
}

func (c *Container) RestartContainer(ctx context.Context, id string) error {
	return c.client.ContainerRestart(ctx, id, container.StopOptions{})
}

func (c *Container) RemoveContainer(ctx context.Context, id string) error {
	return c.client.ContainerRemove(ctx, id, container.RemoveOptions{})
}

func (c *Container) GetContainerLogs(ctx context.Context, id string) (string, error) {
	logs, err := c.client.ContainerLogs(ctx, id, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})
	if err != nil {
		return "", err
	}

	defer logs.Close()

	buf := new(strings.Builder)

	_, err = io.Copy(buf, logs)

	if err != nil {
		return "", err
	}

	return buf.String(), nil

}

func (c *Container) GetContainerLogsStream(ctx context.Context, id string) (io.ReadCloser, error) {
	logs, err := c.client.ContainerLogs(ctx, id, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})
	if err != nil {
		return nil, err
	}

	return logs, nil
}
