package api

import (
	"context"
	"errors"
	"strings"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/v1tbrah/user-service/internal/model"
	"github.com/v1tbrah/user-service/internal/storage"
	"github.com/v1tbrah/user-service/upbapi"
)

func (a *API) CreateCity(ctx context.Context, req *upbapi.CreateCityRequest) (*upbapi.CreateCityResponse, error) {
	reqName := strings.TrimSpace(req.GetName())
	if reqName == "" {
		return nil, status.Error(codes.InvalidArgument, upbapi.ErrEmptyName.Error())
	}

	id, err := a.storage.CreateCity(ctx, model.City{Name: reqName})
	if err != nil {
		if errors.Is(err, storage.ErrCityAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		log.Error().Err(err).Str("name", reqName).Msg("storage.CreateCity")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &upbapi.CreateCityResponse{Id: id}, nil
}
