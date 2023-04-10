package services

import (
	"context"
	"gqlgen-go/graph/db"
	"gqlgen-go/graph/model"
	"log"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type pullRequestService struct {
	exec boil.ContextExecutor
}

func convertPullRequest(pr *db.Pullrequest) *model.PullRequest {
	prURL, err := model.UnmarshalURI(pr.URL)
	if err != nil {
		log.Println("invalid URI", pr.URL)
	}

	return &model.PullRequest{
		ID:          pr.ID,
		BaseRefName: pr.BaseRefName,
		Closed:      (pr.Closed == 1),
		HeadRefName: pr.HeadRefName,
		URL:         prURL,
		Number:      int(pr.Number),
		Repository:  &model.Repository{ID: pr.Repository},
	}
}

func (p *pullRequestService) GetPullRequestByRepoAndNumber(ctx context.Context, repoID string, number int) (*model.PullRequest, error) {
	pr, err := db.Pullrequests(
		qm.Select(
			db.PullrequestColumns.ID,
			db.PullrequestColumns.BaseRefName,
			db.PullrequestColumns.Closed,
			db.PullrequestColumns.HeadRefName,
			db.PullrequestColumns.URL,
			db.PullrequestColumns.Number,
			db.PullrequestColumns.Repository,
		),
		db.PullrequestWhere.Repository.EQ(repoID),
		db.PullrequestWhere.Number.EQ(int64(number)),
	).One(ctx, p.exec)
	if err != nil {
		return nil, err
	}
	return convertPullRequest(pr), nil
}
