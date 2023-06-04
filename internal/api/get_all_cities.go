package api

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.com/pet-pr-social-network/user-service/internal/model"
	"gitlab.com/pet-pr-social-network/user-service/upbapi"
)

func (a *API) GetAllCities(ctx context.Context, req *upbapi.Empty) (*upbapi.GetAllCitiesResponse, error) {
	cities, err := a.storage.GetAllCities(ctx)
	if err != nil {
		log.Error().Err(err).Msg("storage.GetAllCities")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return modelCitiesToProtoGetAllInterestResponse(cities), nil
}

func modelCitiesToProtoGetAllInterestResponse(cities []model.City) *upbapi.GetAllCitiesResponse {
	resp := &upbapi.GetAllCitiesResponse{
		Cities: make([]*upbapi.City, 0, len(cities)),
	}

	for _, city := range cities {
		resp.Cities = append(resp.Cities, &upbapi.City{
			Id:   city.ID,
			Name: city.Name,
		})
	}

	return resp
}
