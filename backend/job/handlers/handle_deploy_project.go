package handlers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null"
	"github.com/hibiken/asynq"
	"github.com/mujhtech/b0/database/models"
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

		serverPort := generatePort()

		existContainerWithPort, err := docker.IsContainerExist(ctx, con.FilterContainerOption{
			Port: serverPort,
		})

		if err != nil {
			return err
		}

		if existContainerWithPort {
			serverPort = generatePort()
		}

		codeGenOption, err := aa.GetLanguageCodeGeneration(project.Language, project.Framework)

		if err != nil {

			sendEvent(ctx, project.ID, sse.EventTypeTaskFailed, AgentData{
				Error: "failed to find supported language option",
			}, event)

			return nil
		}

		codeGenOption.Workflows = endpoint.Workflows
		codeGenOption.FrameworkInsructions = fmt.Sprintf(codeGenOption.FrameworkInsructions, serverPort)

		var code *aa.CodeGeneration

		if endpoint.CodeGeneration != nil && len(endpoint.CodeGeneration.FileContents) > 0 {
			code = endpoint.CodeGeneration
		} else {
			sendEvent(ctx, project.ID, sse.EventTypeTaskUpdate, AgentData{
				Message: "b0 has started generating the code",
			}, event)

			newCode, agentToken, err := agent.CodeGeneration(ctx, project.Description.String, codeGenOption, aa.WithModel(aa.ToModel(project.Model.String)))

			if err != nil {
				sendEvent(ctx, project.ID, sse.EventTypeTaskFailed, AgentData{
					Message: agentToken.Output,
					Error:   err.Error(),
				}, event)

				return err
			}

			// update endpoint
			endpoint.CodeGeneration = newCode

			if err = store.EndpointRepo.UpdateEndpoint(ctx, endpoint.ID, endpoint); err != nil {
				sendEvent(ctx, project.ID, sse.EventTypeTaskFailed, AgentData{
					Message: "b0 failed to update endpoint",
				}, event)

				return nil
			}

			if err = store.AIUsageRepo.CreateAIUsage(ctx, &models.AIUsage{
				ID:          uuid.New().String(),
				ProjectID:   project.ID,
				EndpointID:  null.NewString(endpoint.ID, true),
				OwnerID:     project.OwnerID,
				Model:       project.Model.String,
				UsageType:   "code_generation",
				InputToken:  agentToken.Input,
				OutputToken: agentToken.Output,
			}); err != nil {
				zerolog.Ctx(ctx).Error().Err(err).Msg("failed to create AI usage")
			}

			zerolog.Ctx(ctx).Info().Msgf("code generation: %v", newCode)

			sendEvent(ctx, project.ID, sse.EventTypeTaskUpdate, AgentData{
				Message: "b0 has successfully generated the code",
				Code:    newCode,
			}, event)

			code = newCode
		}

		if code == nil {
			sendEvent(ctx, project.ID, sse.EventTypeTaskFailed, AgentData{
				Message: "b0 failed to find code generation",
			}, event)
			return nil
		}

		isFolderExist, err := checkIfProjectFolderExists(project.OwnerID, project.Slug)

		if err != nil {
			sendEvent(ctx, project.ID, sse.EventTypeTaskFailed, AgentData{
				Message: "b0 failed to check if folder exists",
				Error:   err.Error(),
			}, event)
			return nil
		}

		if !isFolderExist {
			if err = createProjectFolder(project.OwnerID, project.Slug); err != nil {
				sendEvent(ctx, project.ID, sse.EventTypeTaskFailed, AgentData{
					Message: "b0 failed to create folder",
					Error:   err.Error(),
				}, event)
				return nil
			}
		}

		sendEvent(ctx, project.ID, sse.EventTypeTaskUpdate, AgentData{
			Message: "b0 is currently setting up your project...",
		}, event)

		if err = setupProjectContents(project.OwnerID, project.Slug, code); err != nil {

			if err := removeProjectFolder(project.OwnerID, project.Slug); err != nil {
				zerolog.Ctx(ctx).Error().Err(err).Msg("failed to remove project folder")
			}

			sendEvent(ctx, project.ID, sse.EventTypeTaskFailed, AgentData{
				Message: "b0 failed to setup project contents",
				Error:   err.Error(),
			}, event)
			return nil
		}

		// delay 1 seconds
		time.Sleep(1 * time.Second)

		sendEvent(ctx, project.ID, sse.EventTypeTaskUpdate, AgentData{
			Message: "b0 is currently deploying your project...",
		}, event)

		volumeName := fmt.Sprintf("b0-temp-%s-%s", project.OwnerID, project.Slug)

		// check if volume already exists
		if _, err := docker.InspectVolume(ctx, volumeName); err != nil {
			if _, err := docker.CreateVolume(ctx, volumeName); err != nil {
				sendEvent(ctx, project.ID, sse.EventTypeTaskFailed, AgentData{
					Message: "b0 failed to create volume",
					Error:   err.Error(),
				}, event)
				return nil
			}
		}

		serverUrl := fmt.Sprintf("http://%s:%s", "localhost", serverPort)

		// TODO: deploy project to container
		if !project.ContainerID.Valid || project.ContainerID.String == "" {
			// Check if container already exists
			existContainerWithName, err := docker.IsContainerExist(ctx, con.FilterContainerOption{
				Name: project.Slug,
			})

			if err != nil {
				sendEvent(ctx, project.ID, sse.EventTypeTaskFailed, AgentData{
					Message: "b0 failed to check if container exists",
					Error:   err.Error(),
				}, event)
				return nil
			}

			if !existContainerWithName {

				sendEvent(ctx, project.ID, sse.EventTypeTaskUpdate, AgentData{
					Message: "b0 is currently creating a container for your project...",
				}, event)

				// pull image
				if err = docker.PullImage(ctx, codeGenOption.Image); err != nil {
					return err
				}

				isNodeProject := false

				if strings.Contains(project.Language, "Node") {
					isNodeProject = true
				}

				projectCommand := ""

				if isNodeProject {
					projectCommand = fmt.Sprintf(`
						cd /app && \
						NODE_ENV=development %s && \
						%s && \
						NODE_ENV=production %s
					`, strings.Join(code.InstallCommands, " && "), code.BuildCommands, code.RunCommands)
				} else {
					projectCommand = fmt.Sprintf(`
						cd /app && \
						%s && \
						%s && \
						%s
					`, strings.Join(code.InstallCommands, " && "), code.BuildCommands, code.RunCommands)
				}

				commands := []string{"/bin/sh", "-c", projectCommand}

				envs := []string{
					fmt.Sprintf("B0_PORT=%s", serverPort),
				}

				if isNodeProject {
					envs = append(envs, "NODE_ENV=development")
				}

				if code.EnvVars != nil {
					envs = append(envs, code.EnvVars...)
				}

				newContainerID, err := docker.CreateContainer(ctx, con.CreateContainerOption{
					Name:            project.Slug,
					Port:            serverPort,
					Image:           codeGenOption.Image,
					VolumeName:      volumeName,
					HostConfigBinds: []string{fmt.Sprintf("%s:/app", volumeName)},
					Command:         commands,
					WorkingDir:      "/app",
					Env:             envs,
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
					sendEvent(ctx, project.ID, sse.EventTypeTaskFailed, AgentData{
						Message: "b0 failed to create container",
						Error:   err.Error(),
					}, event)
					return nil
				}

				project.ContainerID = null.NewString(newContainerID, true)
				project.Port = null.NewString(serverPort, true)

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
				sendEvent(ctx, project.ID, sse.EventTypeTaskFailed, AgentData{
					Message: "b0 failed to get container",
					Error:   err.Error(),
				}, event)
				return nil
			}

			tar, err := cloneProjectToTar(project.OwnerID, project.Slug)

			if err != nil {
				sendEvent(ctx, project.ID, sse.EventTypeTaskFailed, AgentData{
					Message: "b0 failed to clone project to tar",
					Error:   err.Error(),
				}, event)
				return nil
			}

			if err = docker.CopyFileToContainer(ctx, project.ContainerID.String, tar, "/app"); err != nil {
				sendEvent(ctx, project.ID, sse.EventTypeTaskFailed, AgentData{
					Message: "b0 failed to copy file to container",
					Error:   err.Error(),
				}, event)
				return nil
			}

			zerolog.Ctx(ctx).Info().Msgf("container: %v", container)

			sendEvent(ctx, project.ID, sse.EventTypeTaskUpdate, AgentData{
				Message: "b0 is starting the container for your project...",
			}, event)

			if container.State.Running {
				// restart container
				if err = docker.RestartContainer(ctx, container.ID); err != nil {
					sendEvent(ctx, project.ID, sse.EventTypeTaskFailed, AgentData{
						Message: "b0 failed to restart container",
						Error:   err.Error(),
					}, event)
					return nil
				}
			} else {
				// start container
				if err = docker.StartContainer(ctx, container.ID); err != nil {
					sendEvent(ctx, project.ID, sse.EventTypeTaskFailed, AgentData{
						Message: "b0 failed to start container",
						Error:   err.Error(),
					}, event)
					return nil
				}
			}
		}

		// delay 1 seconds
		time.Sleep(1 * time.Second)

		sendEvent(ctx, project.ID, sse.EventTypeTaskCompleted, AgentData{
			Message: "b0 has successfully deployed your project",
		}, event)

		return nil
	}
}
