package api

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/v1tbrah/user-service/internal/model"
	"github.com/v1tbrah/user-service/upbapi"
)

func (a *API) GetAllInterests(ctx context.Context, req *upbapi.Empty) (*upbapi.GetAllInterestsResponse, error) {
	interests, err := a.storage.GetAllInterests(ctx)
	if err != nil {
		log.Error().Err(err).Msg("storage.GetAllInterests")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return modelInterestsToProtoGetAllInterestResponse(interests), nil
}

func modelInterestsToProtoGetAllInterestResponse(interests []model.Interest) *upbapi.GetAllInterestsResponse {
	resp := &upbapi.GetAllInterestsResponse{
		Interests: make([]*upbapi.Interest, 0, len(interests)),
	}

	for _, interest := range interests {
		resp.Interests = append(resp.Interests, &upbapi.Interest{
			Id:   interest.ID,
			Name: interest.Name,
		})
	}

	return resp
}
