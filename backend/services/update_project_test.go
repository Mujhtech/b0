package services

import (
	"context"
	"testing"

	"github.com/guregu/null"
	"github.com/mujhtech/b0/api/dto"
	"github.com/mujhtech/b0/database/models"
	"github.com/mujhtech/b0/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUpdateProjectService_Run(t *testing.T) {
	type args struct {
		ctx     context.Context
		project *models.Project
		user    *models.User
		dto     *dto.CreateProjectRequestDto
	}

	type testCase struct {
		name    string
		args    args
		mockFn  func(s *UpdateProjectService)
		wantErr error
	}

	tests := []testCase{
		{
			name: "user is owner, should update project",
			args: args{
				ctx:     context.Background(),
				project: &models.Project{ID: "project-id", OwnerID: "user-id", Name: "test project", Description: null.NewString("test description", true)},
				user:    &models.User{ID: "user-id"},
				dto: &dto.CreateProjectRequestDto{
					Name:        "updated project",
					Description: "updated description",
				},
			},
			mockFn: func(s *UpdateProjectService) {
				pr, _ := s.ProjectRepo.(*mocks.MockProjectRepository)
				pr.EXPECT().
					FindProjectByID(gomock.Any(), "project-id").
					Times(1).
					Return(&models.Project{ID: "project-id", OwnerID: "user-id"}, nil)

				pr.EXPECT().
					UpdateProject(gomock.Any(), gomock.Any()).
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

			service := &UpdateProjectService{
				ProjectRepo: mocks.NewMockProjectRepository(ctrl),
				User:        tt.args.user,
				Body:        tt.args.dto,
				ProjectID:   tt.args.project.ID,
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
			}
		})
	}
}
