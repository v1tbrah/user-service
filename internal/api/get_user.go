package api

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/v1tbrah/user-service/internal/model"
	"github.com/v1tbrah/user-service/upbapi"
)

func (a *API) GetUser(ctx context.Context, req *upbapi.GetUserRequest) (*upbapi.GetUserResponse, error) {
	user, err := a.storage.GetUser(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, upbapi.ErrUserNotFoundByID.Error())
		}
		log.Error().Err(err).Str("id", strconv.Itoa(int(req.GetId()))).Msg("storage.GetUser")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return modelUserToProtoGetUserResponse(user), nil
}

func modelUserToProtoGetUserResponse(user model.User) *upbapi.GetUserResponse {
	resp := &upbapi.GetUserResponse{
		Name:        user.Name,
		Surname:     user.Surname,
		InterestsID: user.InterestsID,
	}
	if user.CityID != nil {
		resp.CityID = *user.CityID
	}
	return resp
}
