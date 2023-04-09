package services

import (
	"context"
	"gqlgen-go/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService: &userService{exec: exec},
	}
}

type Services interface {
	UserService
}

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}

type services struct {
	*userService
}
