package handler

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/mujhtech/b0/api/dto"
	"github.com/mujhtech/b0/api/middleware"
	"github.com/mujhtech/b0/database/models"
	"github.com/mujhtech/b0/internal/pkg/agent"
	"github.com/mujhtech/b0/internal/pkg/request"
	"github.com/mujhtech/b0/internal/pkg/response"
	"github.com/mujhtech/b0/services"
)

const (
	EndpointParamId = "endpoint_id"
)

func getEndpointIdFromPath(r *http.Request) (string, error) {
	rawRef, err := pathParamOrError(r, EndpointParamId)
	if err != nil {
		return "", err
	}

	return url.PathUnescape(rawRef)
}

func (h *Handler) GetEndpoints(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	session, ok := middleware.GetAuthSession(ctx)

	if !ok {
		_ = response.Unauthorized(w, r, nil)
		return
	}

	projectId := queryParamOrDefault(r, "project_id", "")

	findEndpointsService := services.FindEndpointsService{
		EndpointRepo: h.store.EndpointRepo,
		User:         session.User,
		Query: dto.GetEndpointQuery{
			ProjectID: projectId,
		},
	}

	endpoints, err := findEndpointsService.Run(ctx)

	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	_ = response.Ok(w, r, "endpoints retrieved", endpoints)
}

func (h *Handler) CreateEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	session, ok := middleware.GetAuthSession(ctx)

	if !ok {
		_ = response.Unauthorized(w, r, nil)
		return
	}

	dst := new(dto.CreateEndpointRequestDto)

	if err := request.ReadBody(r, dst); err != nil {
		_ = response.BadRequest(w, r, err)
		return
	}

	createEndpointService := services.CreateEndpointService{
		Body:         dst,
		ProjectRepo:  h.store.ProjectRepo,
		EndpointRepo: h.store.EndpointRepo,
		User:         session.User,
	}

	endpoint, err := createEndpointService.Run(ctx)

	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	_ = response.Ok(w, r, "endpoint created successfully", endpoint)
}

func (h *Handler) UpdateEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	session, ok := middleware.GetAuthSession(ctx)

	if !ok {
		_ = response.Unauthorized(w, r, nil)
		return
	}

	endpointID, err := getEndpointIdFromPath(r)

	if err != nil {
		_ = response.BadRequest(w, r, err)
		return
	}

	dst := new(dto.CreateEndpointRequestDto)

	if err := request.ReadBody(r, dst); err != nil {
		_ = response.BadRequest(w, r, err)
		return
	}

	updateEndpointService := services.UpdateEndpointService{
		EndpointID:  endpointID,
		ProjectRepo: h.store.ProjectRepo,
		User:        session.User,
		Body:        dst,
	}

	endpoint, err := updateEndpointService.Run(ctx)

	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	_ = response.Ok(w, r, "endpoint updated successfully", endpoint)
}

func (h *Handler) UpdateEndpointWorkflow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	session, ok := middleware.GetAuthSession(ctx)

	if !ok {
		_ = response.Unauthorized(w, r, nil)
		return
	}

	endpointID, err := getEndpointIdFromPath(r)

	if err != nil {
		_ = response.BadRequest(w, r, err)
		return
	}

	dst := new(dto.UpdateEndpointWorkflowRequestDto)

	if err := request.ReadBody(r, dst); err != nil {
		_ = response.BadRequest(w, r, err)
		return
	}

	findEndpointService := services.FindEndpointService{
		EndpointID:   endpointID,
		EndpointRepo: h.store.EndpointRepo,
		User:         session.User,
	}

	endpoint, err := findEndpointService.Run(ctx)

	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	findProjectService := services.FindProjectService{
		ProjectRepo: h.store.ProjectRepo,
		User:        session.User,
		ProjectID:   endpoint.ProjectID,
	}

	project, err := findProjectService.Run(ctx)

	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	if err := h.store.EndpointRepo.UpdateEndpoint(ctx, endpoint.ID, &models.Endpoint{
		Workflows:      dst.Workflows,
		CodeGeneration: &agent.CodeGeneration{},
	}); err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	// remove containers related to the project
	if project.ContainerID.String != "" {

		con, err := h.docker.GetContainer(ctx, project.ContainerID.String)

		if err != nil {
			_ = response.InternalServerError(w, r, err)
			return
		}

		if err := h.docker.RemoveContainer(ctx, con.ID); err != nil {
			_ = response.InternalServerError(w, r, err)
			return
		}

		volumeName := fmt.Sprintf("b0-temp-%s-%s", project.OwnerID, project.Slug)

		if err := h.docker.RemoveVolume(ctx, volumeName); err != nil {
			_ = response.InternalServerError(w, r, err)
			return
		}

	}

	_ = response.Ok(w, r, "endpoint workflow updated successfully", endpoint)
}
