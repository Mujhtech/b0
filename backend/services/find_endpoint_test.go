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

func TestFindEndpointService_Run(t *testing.T) {
	type args struct {
		ctx      context.Context
		endpoint *models.Endpoint
		user     *models.User
	}

	type testCase struct {
		name    string
		args    args
		mockFn  func(s *FindEndpointService)
		wantErr error
	}

	tests := []testCase{
		{
			name: "should find endpoint by id",
			args: args{
				ctx:      context.Background(),
				endpoint: &models.Endpoint{ID: "endpoint-id", OwnerID: "user-id", Name: "test endpoint", Description: null.NewString("test description", true)},
				user:     &models.User{ID: "user-id"},
			},
			mockFn: func(s *FindEndpointService) {
				em, _ := s.EndpointRepo.(*mocks.MockEndpointRepository)
				em.EXPECT().
					FindEndpointByID(gomock.Any(), "endpoint-id").
					Times(1).
					Return(&models.Endpoint{ID: "endpoint-id", OwnerID: "user-id"}, nil)
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := &FindEndpointService{
				EndpointRepo: mocks.NewMockEndpointRepository(ctrl),
				User:         tt.args.user,
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
				require.Equal(t, tt.args.endpoint.ID, endpoint.ID)
			}
		})
	}
}
