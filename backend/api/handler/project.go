package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/mujhtech/b0/api/dto"
	"github.com/mujhtech/b0/api/middleware"
	"github.com/mujhtech/b0/internal/pkg/agent"
	"github.com/mujhtech/b0/internal/pkg/request"
	"github.com/mujhtech/b0/internal/pkg/response"
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

	projectTitleAndSlug, err := h.agent.GenerateTitleAndSlug(ctx, dst.Prompt)

	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	zerolog.Ctx(ctx).Info().Msgf("Generated title and slug: %s", projectTitleAndSlug)

	if strings.HasPrefix(projectTitleAndSlug, "I'm b0, an AI assistant") {
		_ = response.BadRequest(w, r, fmt.Errorf("invalid prompt"))
		return
	}

	// cleanup the message
	projectTitleAndSlug = strings.ReplaceAll(projectTitleAndSlug, "Generated title and slug:", "")
	projectTitleAndSlug = strings.ReplaceAll(projectTitleAndSlug, "json", "")
	projectTitleAndSlug = strings.ReplaceAll(projectTitleAndSlug, "```", "")
	projectTitleAndSlug = strings.ReplaceAll(projectTitleAndSlug, "\n", "")
	projectTitleAndSlug = strings.ReplaceAll(projectTitleAndSlug, `\`, "")

	zerolog.Ctx(ctx).Info().Msgf("cleanup response: %s", strings.TrimSpace(projectTitleAndSlug))

	// unmarshal the projectTitleAndSlug
	var agentProjectTitleAndSlug *agent.ProjectTitleAndSlug

	if err := json.Unmarshal([]byte(projectTitleAndSlug), &agentProjectTitleAndSlug); err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	zerolog.Ctx(ctx).Info().Msgf("%s", agentProjectTitleAndSlug)

	createProjectService := services.CreateProjectService{
		Body:                dst,
		ProjectTitleAndSlug: agentProjectTitleAndSlug,
		ProjectRepo:         h.store.ProjectRepo,
		User:                session.User,
	}

	project, err := createProjectService.Run(ctx)

	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
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
