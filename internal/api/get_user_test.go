package api

import (
	"context"
	"database/sql"
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

func TestAPI_GetUser(t *testing.T) {
	tests := []struct {
		name            string
		mockStorage     func(t *testing.T) *mocks.Storage
		req             *upbapi.GetUserRequest
		expectedResp    *upbapi.GetUserResponse
		wantErr         bool
		expectedErr     error
		expectedErrCode codes.Code
	}{
		{
			name: "OK",
			mockStorage: func(t *testing.T) *mocks.Storage {
				testStorage := mocks.NewStorage(t)
				userFromStorage := model.User{
					ID:          1,
					Name:        "TestName",
					Surname:     "TestSurname",
					InterestsID: []int64{1, 2, 3},
					CityID: func() *int64 {
						var id int64 = 1
						return &id
					}(),
				}
				testStorage.On("GetUser",
					mock.MatchedBy(func(ctx context.Context) bool { return true }), int64(1)).
					Return(userFromStorage, nil).
					Once()
				return testStorage
			},
			req: &upbapi.GetUserRequest{Id: int64(1)},
			expectedResp: &upbapi.GetUserResponse{
				Name:        "TestName",
				Surname:     "TestSurname",
				InterestsID: []int64{1, 2, 3},
				CityID:      1,
			},
		},
		{
			name: "not found",
			mockStorage: func(t *testing.T) *mocks.Storage {
				testStorage := mocks.NewStorage(t)
				testStorage.On("GetUser",
					mock.MatchedBy(func(ctx context.Context) bool { return true }), int64(1)).
					Return(model.User{}, sql.ErrNoRows).
					Once()
				return testStorage
			},
			req:             &upbapi.GetUserRequest{Id: int64(1)},
			wantErr:         true,
			expectedErr:     upbapi.ErrUserNotFoundByID,
			expectedErrCode: codes.NotFound,
		},
		{
			name: "unexpected err on storage.GetUser",
			mockStorage: func(t *testing.T) *mocks.Storage {
				testStorage := mocks.NewStorage(t)
				testStorage.On("GetUser",
					mock.MatchedBy(func(ctx context.Context) bool { return true }), int64(1)).
					Return(model.User{}, errors.New("unexpected err")).
					Once()
				return testStorage
			},
			req:             &upbapi.GetUserRequest{Id: int64(1)},
			wantErr:         true,
			expectedErr:     errors.New("unexpected"),
			expectedErrCode: codes.Internal,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &API{storage: tt.mockStorage(t)}
			resp, err := a.GetUser(context.Background(), tt.req)

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
