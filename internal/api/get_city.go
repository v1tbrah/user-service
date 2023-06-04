package api

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.com/pet-pr-social-network/user-service/internal/model"
	"gitlab.com/pet-pr-social-network/user-service/upbapi"
)

func (a *API) GetCity(ctx context.Context, req *upbapi.GetCityRequest) (*upbapi.GetCityResponse, error) {
	city, err := a.storage.GetCity(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, upbapi.ErrCityNotFoundByID.Error())
		}
		log.Error().Err(err).Str("id", strconv.Itoa(int(req.GetId()))).Msg("storage.GetCity")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return modelICityToProtoGetInterestResponse(city), nil
}

func modelICityToProtoGetInterestResponse(city model.City) *upbapi.GetCityResponse {
	return &upbapi.GetCityResponse{
		City: &upbapi.City{
			Id:   city.ID,
			Name: city.Name,
		},
	}
}
