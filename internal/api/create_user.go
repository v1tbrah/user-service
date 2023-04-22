package api

import (
	"context"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.com/pet-pr-social-network/user-service/internal/model"
	"gitlab.com/pet-pr-social-network/user-service/pbapi"
)

func (a *API) CreateUser(ctx context.Context, req *pbapi.CreateUserRequest) (*pbapi.CreateUserResponse, error) {
	req.Name = strings.TrimSpace(req.GetName())
	if req.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, pbapi.ErrEmptyName.Error())
	}

	req.Surname = strings.TrimSpace(req.GetSurname())
	if req.GetSurname() == "" {
		return nil, status.Error(codes.InvalidArgument, pbapi.ErrEmptySurname.Error())
	}

	id, err := a.storage.CreateUser(ctx, protoCreateUserRequestToModelUser(req))
	if err != nil {
		log.Error().Err(err).Str("user", fmt.Sprintf("%+v", req)).Msg("storage.CreateUser")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pbapi.CreateUserResponse{Id: id}, nil
}

func protoCreateUserRequestToModelUser(req *pbapi.CreateUserRequest) model.User {
	user := model.User{
		Name:        req.GetName(),
		Surname:     req.GetSurname(),
		InterestsID: req.GetInterestsID(),
	}
	cityID := req.GetCityID()
	if cityID != 0 {
		user.CityID = &cityID
	}
	return user
}