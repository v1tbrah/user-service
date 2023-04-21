package api

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.com/pet-pr-social-network/user-service/internal/model"
	"gitlab.com/pet-pr-social-network/user-service/pbapi"
)

func (a *API) GetAllInterests(ctx context.Context, req *pbapi.Empty) (*pbapi.GetAllInterestsResponse, error) {
	interests, err := a.storage.GetAllInterests(ctx)
	if err != nil {
		log.Error().Err(err).Msg("storage.GetAllInterests")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return modelInterestsToProtoGetAllInterestResponse(interests), nil
}

func modelInterestsToProtoGetAllInterestResponse(interests []model.Interest) *pbapi.GetAllInterestsResponse {
	resp := &pbapi.GetAllInterestsResponse{
		Interests: make([]*pbapi.Interest, 0, len(interests)),
	}

	for _, interest := range interests {
		resp.Interests = append(resp.Interests, &pbapi.Interest{
			Id:   interest.ID,
			Name: interest.Name,
		})
	}

	return resp
}
