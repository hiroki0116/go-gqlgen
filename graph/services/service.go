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
		issueService:       &issueService{exec: exec},
		pullRequestService: &pullRequestService{exec: exec},
	}
}

type Services interface {
	UserService
	Reposervice
	ProjectItemService
	IssueService
	PullRequestService
}

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
}

type Reposervice interface {
	GetRepoByFullName(ctx context.Context, name, owner string) (*model.Repository, error)
	GetRepoByID(ctx context.Context, id string) (*model.Repository, error)
}

type ProjectItemService interface {
	AddIssueInProjectV2(ctx context.Context, projectID, issueID string) (*model.ProjectV2Item, error)
}

type IssueService interface {
	GetIssueByRepoAndNumber(ctx context.Context, repoID string, number int) (*model.Issue, error)
	ListIssueInRepository(ctx context.Context, repoID string, after *string, before *string, first *int, last *int) (*model.IssueConnection, error)
}

type PullRequestService interface {
	GetPullRequestByRepoAndNumber(ctx context.Context, repoID string, number int) (*model.PullRequest, error)
	ListPullRequestInRepository(ctx context.Context, repoID string, after *string, before *string, first *int, last *int) (*model.PullRequestConnection, error)
}

type services struct {
	*userService
	*repoService
	*projectItemService
	*issueService
	*pullRequestService
}
