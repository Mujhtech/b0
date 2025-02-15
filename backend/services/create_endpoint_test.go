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

func TestCreateEndpointService_Run(t *testing.T) {
	type args struct {
		ctx      context.Context
		endpoint *models.Endpoint
		project  *models.Project
		user     *models.User
		dto      *dto.CreateEndpointRequestDto
	}

	type testCase struct {
		name    string
		args    args
		mockFn  func(s *CreateEndpointService)
		wantErr error
	}

	tests := []testCase{
		{
			name: "user is owner, should create endpoint",
			args: args{
				ctx:      context.Background(),
				endpoint: &models.Endpoint{ID: "endpoint-id", OwnerID: "user-id", ProjectID: "project-id", Name: "test endpoint", Description: null.NewString("test description", true)},
				project:  &models.Project{ID: "project-id", OwnerID: "user-id"},
				user:     &models.User{ID: "user-id"},
				dto: &dto.CreateEndpointRequestDto{
					Name:        "test endpoint",
					Description: "test description",
					ProjectID:   "project-id", // Add ProjectID to DTO
				},
			},
			mockFn: func(s *CreateEndpointService) {
				pm, _ := s.ProjectRepo.(*mocks.MockProjectRepository)
				pm.EXPECT().
					FindProjectByID(gomock.Any(), "project-id").
					Times(1).
					Return(&models.Project{ID: "project-id", OwnerID: "user-id"}, nil)

				em, _ := s.EndpointRepo.(*mocks.MockEndpointRepository)
				em.EXPECT().
					CreateEndpoint(gomock.Any(), gomock.Any()).
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

			service := &CreateEndpointService{
				ProjectRepo:  mocks.NewMockProjectRepository(ctrl),
				EndpointRepo: mocks.NewMockEndpointRepository(ctrl),
				User:         tt.args.user,
				Body:         tt.args.dto,
			}

			if tt.mockFn != nil {
				tt.mockFn(service)
			}

			endpoint, err := service.Run(tt.args.ctx)

			if tt.wantErr != nil {
				require.Error(t, err)
				require.Nil(t, endpoint)
				require.Equal(t, tt.wantErr, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, endpoint)
				require.Equal(t, tt.args.endpoint.Name, endpoint.Name)
			}
		})
	}
}
