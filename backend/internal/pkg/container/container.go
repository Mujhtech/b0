package container

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/volume"
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

func (c *Container) GetContainer(ctx context.Context, id string) (*container.InspectResponse, error) {
	container, _, err := c.client.ContainerInspectWithRaw(ctx, id, true)
	return &container, err
}

func (c *Container) IsContainerExist(ctx context.Context, opts FilterContainerOption) (bool, error) {

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

func (c *Container) GetContainerList(ctx context.Context, opts FilterContainerOption) ([]container.Summary, error) {
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

func (c *Container) CreateContainer(ctx context.Context, opts CreateContainerOption) (string, error) {
	resp, err := c.client.ContainerCreate(ctx, &container.Config{
		OpenStdin:    true,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
		WorkingDir:   opts.WorkingDir,
		Image:        opts.Image,
		Cmd:          opts.Command,
		Entrypoint:   opts.Entrypoint,
		Env:          opts.Env,
		ExposedPorts: nat.PortSet{
			nat.Port(fmt.Sprintf("%s/tcp", opts.Port)): struct{}{},
		},
		Labels: opts.Labels,
	}, &container.HostConfig{
		Binds: opts.HostConfigBinds,
		PortBindings: map[nat.Port][]nat.PortBinding{
			nat.Port(fmt.Sprintf("%s/tcp", opts.Port)): {
				{
					HostIP:   "0.0.0.0", // Add this to bind to all interfaces
					HostPort: opts.Port,
				},
			},
		},
		Resources: container.Resources{
			CPUQuota:   100000,
			CPUPeriod:  100000,
			Memory:     512 * 1024 * 1024,  // Add memory limit (512MB)
			MemorySwap: 1024 * 1024 * 1024, // Add swap limit (1GB)
		},
		RestartPolicy: container.RestartPolicy{
			Name: "unless-stopped", // Add restart policy
		},
		NetworkMode: "bridge", // Specify network mode
	}, &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			"bridge": {}, // Configure bridge network
		},
	}, nil, opts.Name)
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

func (c *Container) ContainerExec(ctx context.Context, id string, src string) error {

	exec, err := c.client.ContainerExecCreate(ctx, id, container.ExecOptions{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
		Cmd:          []string{"sh", "-c"},
	})
	if err != nil {
		return err
	}

	return c.client.ContainerExecStart(ctx, exec.ID, container.ExecStartOptions{})
}

func (c *Container) CopyFileToContainer(ctx context.Context, id string, src io.Reader, dst string) error {
	return c.client.CopyToContainer(ctx, id, dst, src, container.CopyToContainerOptions{})
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
		Follow:     true,
		Details:    false,
		Timestamps: false,
		Tail:       "all",
	})
	if err != nil {
		return nil, err
	}

	return logs, nil
}

func (c *Container) GetVolume(ctx context.Context, id string) (*volume.Volume, error) {
	volume, _, err := c.client.VolumeInspectWithRaw(ctx, id)
	return &volume, err
}

func (c *Container) CreateVolume(ctx context.Context, name string) (*volume.Volume, error) {
	volume, err := c.client.VolumeCreate(ctx, volume.CreateOptions{
		Name:   name,
		Driver: "local",
	})
	return &volume, err
}

func (c *Container) RemoveVolume(ctx context.Context, id string) error {
	return c.client.VolumeRemove(ctx, id, true)
}

func (c *Container) InspectVolume(ctx context.Context, id string) (*volume.Volume, error) {
	volume, _, err := c.client.VolumeInspectWithRaw(ctx, id)
	return &volume, err
}
