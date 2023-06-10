package api

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/v1tbrah/user-service/internal/api/mocks"
	"github.com/v1tbrah/user-service/internal/model"
	"github.com/v1tbrah/user-service/upbapi"
)

func TestAPI_GetAllCities(t *testing.T) {
	tests := []struct {
		name            string
		mockStorage     func(t *testing.T) *mocks.Storage
		expectedResp    *upbapi.GetAllCitiesResponse
		wantErr         bool
		expectedErr     error
		expectedErrCode codes.Code
	}{
		{
			name: "OK",
			mockStorage: func(t *testing.T) *mocks.Storage {
				testStorage := mocks.NewStorage(t)
				citiesFromStorage := []model.City{
					{
						ID:   1,
						Name: "TestName",
					},
					{
						ID:   2,
						Name: "TestName2",
					},
				}
				testStorage.On("GetAllCities",
					mock.MatchedBy(func(ctx context.Context) bool { return true })).
					Return(citiesFromStorage, nil).
					Once()
				return testStorage
			},
			expectedResp: &upbapi.GetAllCitiesResponse{
				Cities: []*upbapi.City{
					{
						Id:   1,
						Name: "TestName",
					},
					{
						Id:   2,
						Name: "TestName2",
					},
				},
			},
		},
		{
			name: "unexpected err on storage.GetAllCities",
			mockStorage: func(t *testing.T) *mocks.Storage {
				testStorage := mocks.NewStorage(t)
				testStorage.On("GetAllCities",
					mock.MatchedBy(func(ctx context.Context) bool { return true })).
					Return([]model.City{}, errors.New("unexpected err")).
					Once()
				return testStorage
			},
			wantErr:         true,
			expectedErr:     errors.New("unexpected"),
			expectedErrCode: codes.Internal,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &API{storage: tt.mockStorage(t)}
			resp, err := a.GetAllCities(context.Background(), &upbapi.Empty{})

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErrCode, status.Code(err))
				assert.Contains(t, err.Error(), tt.expectedErr.Error())
			}

			if !tt.wantErr {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedResp, resp)
			}
		})
	}
}
