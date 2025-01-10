package services

import (
	"context"
	"testing"

	"github.com/guregu/null"
	"github.com/mujhtech/b0/database/models"
	"github.com/mujhtech/b0/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestFindProjectsService_Run(t *testing.T) {
	type args struct {
		ctx      context.Context
		projects []*models.Project
		user     *models.User
	}

	type testCase struct {
		name    string
		args    args
		mockFn  func(s *FindProjectsService)
		wantErr error
	}

	tests := []testCase{
		{
			name: "should find projects by user",
			args: args{
				ctx: context.Background(),
				projects: []*models.Project{
					{ID: "project-1", OwnerID: "user-id", Name: "test project 1", Description: null.NewString("description 1", true)},
					{ID: "project-2", OwnerID: "user-id", Name: "test project 2", Description: null.NewString("description 2", true)},
				},
				user: &models.User{ID: "user-id"},
			},
			mockFn: func(s *FindProjectsService) {
				pr, _ := s.ProjectRepo.(*mocks.MockProjectRepository)
				pr.EXPECT().
					FindProjectByOwnerID(gomock.Any(), "user-id").
					Times(1).
					Return([]*models.Project{
						{ID: "project-1", OwnerID: "user-id"},
						{ID: "project-2", OwnerID: "user-id"},
					}, nil)
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := &FindProjectsService{
				ProjectRepo: mocks.NewMockProjectRepository(ctrl),
				User:        tt.args.user,
			}

			if tt.mockFn != nil {
				tt.mockFn(service)
			}

			projects, err := service.Run(tt.args.ctx)

			if tt.wantErr != nil {
				require.Error(t, err)
				require.Nil(t, projects)
				require.Equal(t, tt.wantErr, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, projects)
				require.Equal(t, len(tt.args.projects), len(projects))
			}
		})
	}
}
