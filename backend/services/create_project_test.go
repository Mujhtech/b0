package services

import (
	"context"
	"testing"

	"github.com/guregu/null"
	"github.com/mujhtech/b0/api/dto"
	"github.com/mujhtech/b0/database/models"
	"github.com/mujhtech/b0/internal/pkg/agent"
	"github.com/mujhtech/b0/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreateProjectService_Run(t *testing.T) {
	type args struct {
		ctx                 context.Context
		project             *models.Project
		framework           agent.CodeGenerationOption
		user                *models.User
		dto                 *dto.CreateProjectRequestDto
		projectTitleAndSlug *agent.ProjectTitleAndSlug
	}

	type testCase struct {
		name    string
		args    args
		mockFn  func(s *CreateProjectService)
		wantErr error
	}

	tests := []testCase{
		{
			name: "should create project",
			args: args{
				ctx:     context.Background(),
				project: &models.Project{ID: "project-id", OwnerID: "user-id", Name: "test project", Description: null.NewString("test description", true)},
				user:    &models.User{ID: "user-id"},
				dto: &dto.CreateProjectRequestDto{
					Prompt:     "test project",
					IsTemplate: false,
				},
				projectTitleAndSlug: &agent.ProjectTitleAndSlug{
					Title:       "test project",
					Slug:        "test-project",
					Description: "test description",
				},
				framework: agent.CodeGenerationOption{
					Framework: "Chi",
					Language:  "Go",
				},
			},
			mockFn: func(s *CreateProjectService) {
				pr, _ := s.ProjectRepo.(*mocks.MockProjectRepository)
				pr.EXPECT().
					CreateProject(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil)
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := &CreateProjectService{
				ProjectRepo:         mocks.NewMockProjectRepository(ctrl),
				User:                tt.args.user,
				Body:                tt.args.dto,
				ProjectTitleAndSlug: tt.args.projectTitleAndSlug,
				Framework:           tt.args.framework,
			}

			if tt.mockFn != nil {
				tt.mockFn(service)
			}

			project, err := service.Run(tt.args.ctx)

			if tt.wantErr != nil {
				require.Error(t, err)
				require.Nil(t, project)
				require.Equal(t, tt.wantErr, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, project)
				require.Equal(t, tt.args.project.Name, project.Name)
			}
		})
	}
}
