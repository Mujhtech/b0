package handler

import (
	"fmt"
	"net/http"

	"github.com/mujhtech/b0/api/dto"
	"github.com/mujhtech/b0/api/middleware"
	"github.com/mujhtech/b0/internal/pkg/request"
	"github.com/mujhtech/b0/internal/pkg/response"
	"github.com/mujhtech/b0/internal/util"
	jobHandlers "github.com/mujhtech/b0/job/handlers"
	"github.com/mujhtech/b0/services"
)

func (h *Handler) GetScret(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	session, ok := middleware.GetAuthSession(ctx)

	if !ok {
		_ = response.Unauthorized(w, r, nil)
		return
	}

	projectId, err := getProjectIdFromPath(r)

	if err != nil {
		_ = response.BadRequest(w, r, err)
		return
	}

	endpointId := queryParamOrDefault(r, "endpoint", "")

	findProjectService := services.FindProjectService{
		ProjectID:   projectId,
		ProjectRepo: h.store.ProjectRepo,
		User:        session.User,
	}

	project, err := findProjectService.Run(ctx)

	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	if endpointId != "" {

		findEndpointService := services.FindEndpointService{
			EndpointID:   endpointId,
			EndpointRepo: h.store.EndpointRepo,
			User:         session.User,
		}

		endpoint, err := findEndpointService.Run(ctx)

		if err != nil {
			_ = response.InternalServerError(w, r, err)
			return
		}

		endpointId = endpoint.ID

	}

	secrets, err := jobHandlers.GetEnvVars(ctx, h.secretManager, project.ID, endpointId)

	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	_ = response.Ok(w, r, "environmental variable retrieved", secrets)
}

func (h *Handler) CreateOrUpdateScret(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	session, ok := middleware.GetAuthSession(ctx)

	if !ok {
		_ = response.Unauthorized(w, r, nil)
		return
	}

	projectId, err := getProjectIdFromPath(r)

	if err != nil {
		_ = response.BadRequest(w, r, err)
		return
	}

	endpointId := queryParamOrDefault(r, "endpoint", "")

	dst := new(dto.SecretRequestDto)

	if err := request.ReadBody(r, dst); err != nil {
		_ = response.BadRequest(w, r, err)
		return
	}

	findProjectService := services.FindProjectService{
		ProjectID:   projectId,
		ProjectRepo: h.store.ProjectRepo,
		User:        session.User,
	}

	project, err := findProjectService.Run(ctx)

	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	secretId := project.ID

	if endpointId != "" {

		findEndpointService := services.FindEndpointService{
			EndpointID:   endpointId,
			EndpointRepo: h.store.EndpointRepo,
			User:         session.User,
		}

		endpoint, err := findEndpointService.Run(ctx)

		if err != nil {
			_ = response.InternalServerError(w, r, err)
			return
		}

		secretId = fmt.Sprintf("%s_%s", secretId, endpoint.ID)

	}

	secretName := fmt.Sprintf("projects/b0/%s/env-variables", secretId)

	rawSecret, err := util.MarshalJSON(dst.Secrets)
	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	err = h.secretManager.SetSecret(ctx, secretName, rawSecret)

	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	_ = response.Ok(w, r, "environmental variable saved", nil)
}
