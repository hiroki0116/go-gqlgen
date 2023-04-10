package services

import (
	"context"
	"gqlgen-go/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService:        &userService{exec: exec},
		repoService:        &repoService{exec: exec},
		projectItemService: &projectItemService{exec: exec},
	}
}

type Services interface {
	UserService
	Reposervice
	ProjectItemService
}

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}

type Reposervice interface {
	GetRepoByFullName(ctx context.Context, name, owner string) (*model.Repository, error)
}

type ProjectItemService interface {
	AddIssueInProjectV2(ctx context.Context, projectID, issueID string) (*model.ProjectV2Item, error)
}

type services struct {
	*userService
	*repoService
	*projectItemService
}
