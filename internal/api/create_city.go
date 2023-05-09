package api

import (
	"context"
	"errors"
	"strings"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.com/pet-pr-social-network/user-service/internal/model"
	"gitlab.com/pet-pr-social-network/user-service/internal/storage"
	"gitlab.com/pet-pr-social-network/user-service/pbapi"
)

func (a *API) CreateCity(ctx context.Context, req *pbapi.CreateCityRequest) (*pbapi.CreateCityResponse, error) {
	reqName := strings.TrimSpace(req.GetName())
	if reqName == "" {
		return nil, status.Error(codes.InvalidArgument, pbapi.ErrEmptyName.Error())
	}

	id, err := a.storage.CreateCity(ctx, model.City{Name: reqName})
	if err != nil {
		if errors.Is(err, storage.ErrCityAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		log.Error().Err(err).Str("name", reqName).Msg("storage.CreateCity")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pbapi.CreateCityResponse{Id: id}, nil
}
