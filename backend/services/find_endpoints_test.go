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

func TestFindEndpointsService_Run(t *testing.T) {
	type args struct {
		ctx       context.Context
		endpoints []*models.Endpoint
		user      *models.User
		projectID string
	}

	type testCase struct {
		name    string
		args    args
		mockFn  func(s *FindEndpointsService)
		wantErr error
	}

	tests := []testCase{
		{
			name: "should find endpoints by project id",
			args: args{
				ctx: context.Background(),
				endpoints: []*models.Endpoint{
					{ID: "endpoint-1", OwnerID: "user-id", ProjectID: "project-id", Name: "test endpoint 1", Description: null.NewString("description 1", true)},
					{ID: "endpoint-2", OwnerID: "user-id", ProjectID: "project-id", Name: "test endpoint 2", Description: null.NewString("description 2", true)},
				},
				user:      &models.User{ID: "user-id"},
				projectID: "project-id",
			},
			mockFn: func(s *FindEndpointsService) {
				em, _ := s.EndpointRepo.(*mocks.MockEndpointRepository)
				em.EXPECT().
					FindEndpointByProjectID(gomock.Any(), "project-id").
					Times(1).
					Return([]*models.Endpoint{
						{ID: "endpoint-1", OwnerID: "user-id", ProjectID: "project-id"},
						{ID: "endpoint-2", OwnerID: "user-id", ProjectID: "project-id"},
					}, nil)
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := &FindEndpointsService{
				EndpointRepo: mocks.NewMockEndpointRepository(ctrl),
				User:         tt.args.user,
				Query: dto.GetEndpointQuery{
					ProjectID: tt.args.projectID,
				},
			}

			if tt.mockFn != nil {
				tt.mockFn(service)
			}

			endpoints, err := service.Run(tt.args.ctx)

			if tt.wantErr != nil {
				require.Error(t, err)
				require.Nil(t, endpoints)
				require.Equal(t, tt.wantErr, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, endpoints)
				require.Equal(t, len(tt.args.endpoints), len(endpoints))
			}
		})
	}
}
