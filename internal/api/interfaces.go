package api

import (
	"context"

	"github.com/v1tbrah/user-service/internal/model"
)

//go:generate mockery --name Storage
type Storage interface {
	CreateUser(ctx context.Context, user model.User) (id int64, err error)
	GetUser(ctx context.Context, id int64) (user model.User, err error)

	CreateInterest(ctx context.Context, interest model.Interest) (id int64, err error)
	GetInterest(ctx context.Context, id int64) (interest model.Interest, err error)
	GetAllInterests(ctx context.Context) (interests []model.Interest, err error)
	// TODO GetInterestsByUser

	CreateCity(ctx context.Context, city model.City) (id int64, err error)
	GetCity(ctx context.Context, id int64) (city model.City, err error)
	GetAllCities(ctx context.Context) (cities []model.City, err error)
}
