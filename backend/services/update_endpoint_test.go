package services

import (
	"context"
	"testing"

	"github.com/guregu/null"
	"github.com/mujhtech/b0/api/dto"
	"github.com/mujhtech/b0/database/models"
	"github.com/mujhtech/b0/errors"
	"github.com/mujhtech/b0/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUpdateEndpointService_Run(t *testing.T) {
	type args struct {
		ctx      context.Context
		endpoint *models.Endpoint
		user     *models.User
		project  *models.Project
		dto      *dto.CreateEndpointRequestDto
	}

	type testCase struct {
		name    string
		args    args
		mockFn  func(s *UpdateEndpointService)
		wantErr error
	}

	tests := []testCase{
		{
			name: "user is owner, should update endpoint",
			args: args{
				ctx: context.Background(),
				endpoint: &models.Endpoint{
					ID:          "endpoint-id",
					OwnerID:     "user-id",
					ProjectID:   "project-id",
					Name:        "test endpoint",
					Description: null.NewString("test description", true),
				},
				project: &models.Project{
					ID:      "project-id",
					OwnerID: "user-id",
				},
				user: &models.User{
					ID: "user-id",
				},
				dto: &dto.CreateEndpointRequestDto{
					Name:        "updated endpoint",
					Description: "updated description",
					ProjectID:   "project-id",
				},
			},
			mockFn: func(s *UpdateEndpointService) {
				em := s.EndpointRepo.(*mocks.MockEndpointRepository)
				pm := s.ProjectRepo.(*mocks.MockProjectRepository)

				// First expect FindEndpointByID
				em.EXPECT().
					FindEndpointByID(gomock.Any(), "endpoint-id").
					Times(1).
					Return(&models.Endpoint{
						ID:        "endpoint-id",
						ProjectID: "project-id",
						OwnerID:   "user-id",
					}, nil)

				// Then expect FindProjectByID
				pm.EXPECT().
					FindProjectByID(gomock.Any(), "project-id").
					Times(1).
					Return(&models.Project{
						ID:      "project-id",
						OwnerID: "user-id",
					}, nil)

				// Finally expect UpdateEndpoint
				em.EXPECT().
					UpdateEndpoint(gomock.Any(), &models.Endpoint{
						ID:          "endpoint-id",
						OwnerID:     "user-id",
						ProjectID:   "project-id",
						Name:        "updated endpoint",
						Slug:        "updated-endpoint",
						Description: null.NewString("updated description", true),
					}).
					Times(1).
					Return(nil)
			},
			wantErr: nil,
		},
		{
			name: "unauthorized - different owner",
			args: args{
				ctx: context.Background(),
				endpoint: &models.Endpoint{
					ID:          "endpoint-id",
					OwnerID:     "other-user-id",
					ProjectID:   "project-id",
					Name:        "test endpoint",
					Description: null.NewString("test description", true),
				},
				user: &models.User{
					ID: "user-id",
				},
				dto: &dto.CreateEndpointRequestDto{
					Name:        "updated endpoint",
					Description: "updated description",
					ProjectID:   "project-id",
				},
			},
			mockFn: func(s *UpdateEndpointService) {
				em := s.EndpointRepo.(*mocks.MockEndpointRepository)
				pm := s.ProjectRepo.(*mocks.MockProjectRepository)

				em.EXPECT().
					FindEndpointByID(gomock.Any(), "endpoint-id").
					Times(1).
					Return(&models.Endpoint{
						ID:        "endpoint-id",
						ProjectID: "project-id",
						OwnerID:   "other-user-id",
					}, nil)

				pm.EXPECT().
					FindProjectByID(gomock.Any(), "project-id").
					Times(1).
					Return(&models.Project{
						ID:      "project-id",
						OwnerID: "user-id",
					}, nil)
				// Remove the UpdateEndpoint expectation since it should fail before that
			},
			wantErr: errors.ErrNotAuthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := &UpdateEndpointService{
				EndpointRepo: mocks.NewMockEndpointRepository(ctrl),
				ProjectRepo:  mocks.NewMockProjectRepository(ctrl),
				User:         tt.args.user,
				Body:         tt.args.dto,
				EndpointID:   tt.args.endpoint.ID,
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
			}
		})
	}
}
