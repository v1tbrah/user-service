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

	"gitlab.com/pet-pr-social-network/user-service/internal/api/mocks"
	"gitlab.com/pet-pr-social-network/user-service/internal/model"
	"gitlab.com/pet-pr-social-network/user-service/pbapi"
)

func TestAPI_GetInterest(t *testing.T) {
	tests := []struct {
		name            string
		mockStorage     func(t *testing.T) *mocks.Storage
		req             *pbapi.GetInterestRequest
		expectedResp    *pbapi.GetInterestResponse
		wantErr         bool
		expectedErr     error
		expectedErrCode codes.Code
	}{
		{
			name: "OK",
			mockStorage: func(t *testing.T) *mocks.Storage {
				testStorage := mocks.NewStorage(t)
				interestFromStorage := model.Interest{
					ID:   1,
					Name: "TestName",
				}
				testStorage.On("GetInterest",
					mock.MatchedBy(func(ctx context.Context) bool { return true }), int64(1)).
					Return(interestFromStorage, nil).
					Once()
				return testStorage
			},
			req: &pbapi.GetInterestRequest{Id: int64(1)},
			expectedResp: &pbapi.GetInterestResponse{
				Interest: &pbapi.Interest{
					Id:   int64(1),
					Name: "TestName",
				},
			},
		},
		{
			name: "not found",
			mockStorage: func(t *testing.T) *mocks.Storage {
				testStorage := mocks.NewStorage(t)
				testStorage.On("GetInterest",
					mock.MatchedBy(func(ctx context.Context) bool { return true }), int64(1)).
					Return(model.Interest{}, sql.ErrNoRows).
					Once()
				return testStorage
			},
			req:             &pbapi.GetInterestRequest{Id: int64(1)},
			wantErr:         true,
			expectedErr:     pbapi.ErrInterestNotFoundByID,
			expectedErrCode: codes.NotFound,
		},
		{
			name: "unexpected err on storage.GetInterest",
			mockStorage: func(t *testing.T) *mocks.Storage {
				testStorage := mocks.NewStorage(t)
				testStorage.On("GetInterest",
					mock.MatchedBy(func(ctx context.Context) bool { return true }), int64(1)).
					Return(model.Interest{}, errors.New("unexpected err")).
					Once()
				return testStorage
			},
			req:             &pbapi.GetInterestRequest{Id: int64(1)},
			wantErr:         true,
			expectedErr:     errors.New("unexpected"),
			expectedErrCode: codes.Internal,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &API{storage: tt.mockStorage(t)}
			resp, err := a.GetInterest(context.Background(), tt.req)

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
