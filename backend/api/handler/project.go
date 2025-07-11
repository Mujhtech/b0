package handler

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/uuid"
	"github.com/mujhtech/b0/api/dto"
	"github.com/mujhtech/b0/api/middleware"
	"github.com/mujhtech/b0/database/models"
	"github.com/mujhtech/b0/database/store"
	"github.com/mujhtech/b0/internal/pkg/agent"
	"github.com/mujhtech/b0/internal/pkg/request"
	"github.com/mujhtech/b0/internal/pkg/response"
	"github.com/mujhtech/b0/internal/util"
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

	// Check project limits based on subscription plan
	count, err := h.store.ProjectRepo.CountByOwnerID(ctx, session.User.ID)
	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	// Check project limits based on plan
	switch session.User.SubscriptionPlan {
	case "free":
		if count >= 3 {
			_ = response.BadRequest(w, r, fmt.Errorf("you have reached the maximum number of projects allowed for the free plan"))
			return
		}
	case "starter":
		if count >= 10 {
			_ = response.BadRequest(w, r, fmt.Errorf("you have reached the maximum number of projects allowed for the starter plan"))
			return
		}
	}

	var framework agent.CodeGenerationOption

	var catalog agent.ModeCatalog

	if dst.Model != "" {
		var agentModelErr error
		catalog, agentModelErr := agent.GetModelCatalog(dst.Model)
		if agentModelErr != nil {
			_ = response.BadRequest(w, r, agentModelErr)
			return
		}

		if (session.User.SubscriptionPlan == "free" || session.User.SubscriptionPlan == "starter") && catalog.IsPremium {
			_ = response.BadRequest(w, r, fmt.Errorf("you are not allowed to use premium models with the free plan"))
			return
		}

		usageCount, err := h.store.AIUsageRepo.GetTotalUsage(ctx, store.TotalAIUsageFilter{
			OwnerID: session.User.ID,
			Range:   store.TotalAIUsageFilterRangeMonth,
		})

		if err != nil {
			_ = response.InternalServerError(w, r, err)
			return
		}

		if session.User.SubscriptionPlan == "premium" && usageCount.TotalUsage >= 20 {
			_ = response.BadRequest(w, r, fmt.Errorf("you have reached the maximum number of requests for the current month"))
			return
		}

		if session.User.SubscriptionPlan == "pro" && usageCount.TotalUsage >= 100 {
			_ = response.BadRequest(w, r, fmt.Errorf("you have reached the maximum number of requests for the current month"))
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

	agentProjectTitleAndSlug, agentToken, err := h.agent.GenerateTitleAndSlug(ctx, dst.Prompt, agent.WithModel(catalog.Model))

	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	// check for duplicate slug
	duplicateProject, err := h.store.ProjectRepo.FindProjectBySlug(ctx, agentProjectTitleAndSlug.Slug)

	if err != nil && !errors.Is(err, store.ErrNotFound) {
		_ = response.BadRequest(w, r, err)
		return
	}

	if duplicateProject != nil {
		// remove last 6 characters from the slug
		newSlug := agentProjectTitleAndSlug.Slug[:len(agentProjectTitleAndSlug.Slug)-6]
		newSlug, _ = util.GeneratePrefixedID(newSlug, "-", 6)

		agentProjectTitleAndSlug.Slug = newSlug
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

	if err = h.store.AIUsageRepo.CreateAIUsage(ctx, &models.AIUsage{
		ID:          uuid.New().String(),
		ProjectID:   project.ID,
		OwnerID:     project.OwnerID,
		Model:       project.Model.String,
		UsageType:   "project_creation",
		InputToken:  agentToken.Input,
		OutputToken: agentToken.Output,
		IsPremium:   catalog.IsPremium,
	}); err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("failed to create AI usage")
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

func (h *Handler) DeleteProject(w http.ResponseWriter, r *http.Request) {
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

	dst := new(dto.DeleteProjectRequestDto)

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

	if dst.Name != project.Name {
		_ = response.BadRequest(w, r, fmt.Errorf("invalid project name"))
		return
	}

	// delete project related resources
	if project.ContainerID.Valid && project.ContainerID.String != "" {
		if err := h.store.ProjectRepo.UpdateProject(ctx, project); err != nil {
			_ = response.InternalServerError(w, r, err)
			return
		}

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

	// delete project
	err = h.store.ProjectRepo.DeleteProject(ctx, project.ID)

	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	_ = response.Ok(w, r, "project deleted successfully", project)
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
