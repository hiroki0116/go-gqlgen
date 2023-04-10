package services

import (
	"context"
	"gqlgen-go/graph/db"
	"gqlgen-go/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type repoService struct {
	exec boil.ContextExecutor
}

func convertRepository(repo *db.Repository) *model.Repository {
	return &model.Repository{
		ID:        repo.ID,
		Name:      repo.Name,
		Owner:     &model.User{ID: repo.Owner},
		CreatedAt: repo.CreatedAt,
	}
}

func (r *repoService) GetRepoByFullName(ctx context.Context, name, owner string) (*model.Repository, error) {
	repo, err := db.Repositories(
		qm.Select(
			db.RepositoryColumns.ID,
			db.RepositoryColumns.Name,
			db.RepositoryColumns.Owner,
			db.RepositoryColumns.CreatedAt,
		),
		db.RepositoryWhere.Owner.EQ(owner),
		db.RepositoryWhere.Name.EQ(name),
	).One(ctx, r.exec)
	if err != nil {
		return nil, err
	}
	return convertRepository(repo), nil
}
