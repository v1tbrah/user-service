package api

import (
	"context"
	"errors"
	"strings"

	"github.com/rs/zerolog/log"
	"gitlab.com/pet-pr-social-network/user-service/internal/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.com/pet-pr-social-network/user-service/internal/model"
	"gitlab.com/pet-pr-social-network/user-service/pbapi"
)

func (a *API) CreateInterest(ctx context.Context, req *pbapi.CreateInterestRequest) (*pbapi.CreateInterestResponse, error) {
	reqName := strings.TrimSpace(req.GetName())
	if reqName == "" {
		return nil, status.Error(codes.InvalidArgument, pbapi.ErrEmptyName.Error())
	}

	id, err := a.storage.CreateInterest(ctx, model.Interest{Name: reqName})
	if err != nil {
		if errors.Is(err, storage.ErrInterestAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		log.Error().Err(err).Str("name", reqName).Msg("storage.CreateInterest")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pbapi.CreateInterestResponse{Id: id}, nil
}
