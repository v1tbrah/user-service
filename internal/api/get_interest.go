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

func (a *API) GetInterest(ctx context.Context, req *upbapi.GetInterestRequest) (*upbapi.GetInterestResponse, error) {
	interest, err := a.storage.GetInterest(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, upbapi.ErrInterestNotFoundByID.Error())
		}
		log.Error().Err(err).Str("id", strconv.Itoa(int(req.GetId()))).Msg("storage.GetInterest")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return modelInterestToProtoGetInterestResponse(interest), nil
}

func modelInterestToProtoGetInterestResponse(interest model.Interest) *upbapi.GetInterestResponse {
	return &upbapi.GetInterestResponse{
		Interest: &upbapi.Interest{
			Id:   interest.ID,
			Name: interest.Name,
		},
	}
}
