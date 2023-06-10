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
	"github.com/v1tbrah/user-service/upbapi"
)

func TestAPI_CreateInterest(t *testing.T) {
	tests := []struct {
		name            string
		mockStorage     func(t *testing.T) *mocks.Storage
		req             *upbapi.CreateInterestRequest
		expectedResp    *upbapi.CreateInterestResponse
		wantErr         bool
		expectedErr     error
		expectedErrCode codes.Code
	}{
		{
			name: "OK",
			mockStorage: func(t *testing.T) *mocks.Storage {
				testStorage := mocks.NewStorage(t)
				testStorage.On("CreateInterest",
					mock.MatchedBy(func(ctx context.Context) bool { return true }), mock.AnythingOfType("model.Interest")).
					Return(int64(1), nil).
					Once()
				return testStorage
			},
			req: &upbapi.CreateInterestRequest{
				Name: "TestName",
			},
			expectedResp: &upbapi.CreateInterestResponse{Id: int64(1)},
		},
		{
			name: "empty name",
			mockStorage: func(t *testing.T) *mocks.Storage {
				return mocks.NewStorage(t)
			},
			req: &upbapi.CreateInterestRequest{
				Name: "",
			},
			wantErr:         true,
			expectedErr:     upbapi.ErrEmptyName,
			expectedErrCode: codes.InvalidArgument,
		},
		{
			name: "empty name with spaces",
			mockStorage: func(t *testing.T) *mocks.Storage {
				return mocks.NewStorage(t)
			},
			req: &upbapi.CreateInterestRequest{
				Name: "   ",
			},
			wantErr:         true,
			expectedErr:     upbapi.ErrEmptyName,
			expectedErrCode: codes.InvalidArgument,
		},
		{
			name: "unexpected err on storage.CreateInterest",
			mockStorage: func(t *testing.T) *mocks.Storage {
				testStorage := mocks.NewStorage(t)
				testStorage.On("CreateInterest",
					mock.MatchedBy(func(ctx context.Context) bool { return true }), mock.AnythingOfType("model.Interest")).
					Return(int64(0), errors.New("unexpected err")).
					Once()
				return testStorage
			},
			req: &upbapi.CreateInterestRequest{
				Name: "TestName",
			},
			wantErr:         true,
			expectedErr:     errors.New("unexpected"),
			expectedErrCode: codes.Internal,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &API{storage: tt.mockStorage(t)}
			resp, err := a.CreateInterest(context.Background(), tt.req)

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
