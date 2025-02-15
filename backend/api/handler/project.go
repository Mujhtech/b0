package handler

import (
	"net/http"
	"net/url"

	"github.com/mujhtech/b0/api/dto"
	"github.com/mujhtech/b0/api/middleware"
	"github.com/mujhtech/b0/internal/pkg/agent"
	"github.com/mujhtech/b0/internal/pkg/request"
	"github.com/mujhtech/b0/internal/pkg/response"
	"github.com/mujhtech/b0/job"
	"github.com/mujhtech/b0/services"
	"github.com/rs/zerolog"
)

const (
	ProjectParamId = "project_id"
)

func getProjectIdFromPath(r *http.Request) (string, error) {
	rawRef, err := pathParamOrError(r, ProjectParamId)
	if err != nil {
		return "", err
	}

	return url.PathUnescape(rawRef)
}

func (h *Handler) GetProjects(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	session, ok := middleware.GetAuthSession(ctx)

	if !ok {
		_ = response.Unauthorized(w, r, nil)
		return
	}

	findProjectsService := services.FindProjectsService{
		ProjectRepo: h.store.ProjectRepo,
		User:        session.User,
	}

	projects, err := findProjectsService.Run(ctx)

	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	_ = response.Ok(w, r, "projects retrieved", projects)
}

func (h *Handler) GetProject(w http.ResponseWriter, r *http.Request) {
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

	findProjectService := services.FindProjectBySlugService{
		Slug:        projectId,
		ProjectRepo: h.store.ProjectRepo,
		User:        session.User,
	}

	project, err := findProjectService.Run(ctx)

	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	_ = response.Ok(w, r, "project retrieved", project)
}

func (h *Handler) CreateProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	session, ok := middleware.GetAuthSession(ctx)

	if !ok {
		_ = response.Unauthorized(w, r, nil)
		return
	}

	dst := new(dto.CreateProjectRequestDto)

	if err := request.ReadBody(r, dst); err != nil {
		_ = response.BadRequest(w, r, err)
		return
	}

	var framework agent.CodeGenerationOption

	var agentModel agent.AgentModel

	if dst.Model != "" {
		var agentModelErr error
		agentModel, agentModelErr = agent.GetModel(dst.Model)
		if agentModelErr != nil {
			_ = response.BadRequest(w, r, agentModelErr)
			return
		}
	}

	if dst.FramekworkID != "" {
		var err error
		framework, err = agent.GetLanguageCodeGenerationByID(dst.FramekworkID)
		if err != nil {
			_ = response.BadRequest(w, r, err)
			return
		}
	}

	agentProjectTitleAndSlug, _, err := h.agent.GenerateTitleAndSlug(ctx, dst.Prompt, agent.WithModel(agentModel))

	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	createProjectService := services.CreateProjectService{
		Body:                dst,
		ProjectTitleAndSlug: agentProjectTitleAndSlug,
		ProjectRepo:         h.store.ProjectRepo,
		User:                session.User,
		Framework:           framework,
	}

	project, err := createProjectService.Run(ctx)

	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	if err = h.job.Client.Enqueue(job.QueueNameDefault, job.JobNameWorkflowCreate, &job.ClientPayload{
		Data: []byte(project.ID),
	}); err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("failed to enqueue job")
	}

	_ = response.Ok(w, r, "project created successfully", project)
}

func (h *Handler) UpdateProject(w http.ResponseWriter, r *http.Request) {
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

	dst := new(dto.CreateProjectRequestDto)

	if err := request.ReadBody(r, dst); err != nil {
		_ = response.BadRequest(w, r, err)
		return
	}

	updateProjectService := services.UpdateProjectService{
		ProjectID:   projectId,
		ProjectRepo: h.store.ProjectRepo,
		User:        session.User,
		Body:        dst,
	}

	project, err := updateProjectService.Run(ctx)

	if err != nil {

		_ = response.InternalServerError(w, r, err)
		return
	}

	_ = response.Ok(w, r, "project updated successfully", project)
}

func (h *Handler) ProjectAction(w http.ResponseWriter, r *http.Request) {
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

	dst := new(dto.ProjectActionRequestDto)

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

	switch dst.Action {
	case "deploy":

		if err = h.job.Client.Enqueue(job.QueueNameDefault, job.JobNameProjectDeploy, &job.ClientPayload{
			Data: []byte(project.ID),
		}); err != nil {
			zerolog.Ctx(ctx).Error().Err(err).Msg("failed to enqueue job")
		}

	case "export":

		if err = h.job.Client.Enqueue(job.QueueNameDefault, job.JobNameProjectExport, &job.ClientPayload{
			Data: []byte(project.ID),
		}); err != nil {
			zerolog.Ctx(ctx).Error().Err(err).Msg("failed to enqueue job")
		}
	default:
		_ = response.BadRequest(w, r, nil)
		return
	}

	_ = response.Ok(w, r, "ok", nil)
}
