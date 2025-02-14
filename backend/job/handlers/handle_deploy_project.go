package handlers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/guregu/null"
	"github.com/hibiken/asynq"
	"github.com/mujhtech/b0/database/store"
	aa "github.com/mujhtech/b0/internal/pkg/agent"
	con "github.com/mujhtech/b0/internal/pkg/container"
	"github.com/mujhtech/b0/internal/pkg/encrypt"
	"github.com/mujhtech/b0/internal/pkg/sse"
	"github.com/rs/zerolog"
)

func HandleDeployProject(aesCfb encrypt.Encrypt, store *store.Store, agent *aa.Agent, event sse.Streamer, docker *con.Container) func(context.Context, *asynq.Task) error {
	return func(ctx context.Context, t *asynq.Task) error {

		projectId, err := aesCfb.Decrypt(string(t.Payload()))

		if err != nil {
			return err
		}

		project, err := store.ProjectRepo.FindProjectByID(ctx, projectId)

		if err != nil {
			return err
		}

		sendEvent(ctx, project.ID, sse.EventTypeTaskStarted, AgentData{
			Message: "b0 is working on your request...",
		}, event)

		endpoints, err := store.EndpointRepo.FindEndpointByProjectID(ctx, project.ID)

		if err != nil {
			return err
		}

		if len(endpoints) == 0 {
			return fmt.Errorf("no endpoints found for project: %s", project.ID)
		}

		endpoint := endpoints[0]

		// delay 1 seconds
		time.Sleep(1 * time.Second)

		sendEvent(ctx, project.ID, sse.EventTypeTaskUpdate, AgentData{
			Message: "b0 has started generating the code",
		}, event)

		code, agentToken, err := agent.CodeGeneration(ctx, project.Description.String, "Go", endpoint.Workflows, aa.WithModel(aa.ToModel(project.Model.String)))

		if err != nil {
			sendEvent(ctx, project.ID, sse.EventTypeTaskFailed, AgentData{
				Message: agentToken.Output,
				Error:   err.Error(),
			}, event)

			return err
		}

		sendEvent(ctx, project.ID, sse.EventTypeTaskUpdate, AgentData{
			Message: "b0 has successfully generated the code",
		}, event)

		zerolog.Ctx(ctx).Info().Msgf("code generation: %v", code)

		// delay 1 seconds
		time.Sleep(1 * time.Second)

		sendEvent(ctx, project.ID, sse.EventTypeTaskUpdate, AgentData{
			Message: "b0 is currently deploying your project...",
		}, event)

		serverUrl := fmt.Sprintf("https://%s.%s", project.Slug, strings.ToLower("b0.dev"))

		// TODO: deploy project to container
		if !project.ContainerID.Valid || project.ContainerID.String == "" {
			// Check if container already exists
			existContainerWithName, err := docker.IsContainerExist(ctx, con.FilterContainerOption{
				Name: project.Slug,
			})

			if err != nil {
				return err
			}

			// check if container with same port exists
			// existContainerWithPort, err := container.IsContainerExist(ctx, con.FilterContainerOption{
			// 	Port: "8080",
			// })

			// if err != nil {
			// 	return err
			// }

			if !existContainerWithName {

				sendEvent(ctx, project.ID, sse.EventTypeTaskUpdate, AgentData{
					Message: "b0 is currently creating a container for your project...",
				}, event)

				image := "golang:1.23-alpine3.20"

				// pull image
				if err = docker.PullImage(ctx, image); err != nil {
					return err
				}

				newContainerID, err := docker.CreateContainer(ctx, con.CreateContainerOption{
					Name:  project.Slug,
					Port:  "8080",
					Image: image,
					Labels: map[string]string{
						"traefik.enable": "true",
						fmt.Sprintf("traefik.http.routers.%s.rule", project.Slug):        fmt.Sprintf("Host(`%s`)", strings.Replace(serverUrl, "https://", "", 1)),
						fmt.Sprintf("traefik.http.routers.%s.entrypoints", project.Slug): "websecure",
						fmt.Sprintf("traefik.http.routers.%s.tls", project.Slug):         "true",
						"project_id":   project.ID,
						"project_name": project.Name,
					},
				})

				if err != nil {
					return err
				}

				project.ContainerID = null.NewString(newContainerID, true)

				// update project
				if err = store.ProjectRepo.UpdateProject(ctx, project); err != nil {
					return err
				}

				sendEvent(ctx, project.ID, sse.EventTypeTaskUpdate, AgentData{
					Message: "b0 has successfully created a container for your project",
				}, event)
			}

		}

		if project.ContainerID.Valid && project.ContainerID.String != "" {
			container, err := docker.GetContainer(ctx, project.ContainerID.String)

			if err != nil {
				return err
			}

			zerolog.Ctx(ctx).Info().Msgf("container: %v", container)

			if container.State.Running {
				// restart container
				if err = docker.RestartContainer(ctx, container.ID); err != nil {
					return err
				}
			}
		}

		// delay 1 seconds
		time.Sleep(1 * time.Second)

		sendEvent(ctx, project.ID, sse.EventTypeTaskCompleted, AgentData{
			Message: "b0 has successfully deployed your project",
			Code:    code,
		}, event)

		return nil
	}
}
