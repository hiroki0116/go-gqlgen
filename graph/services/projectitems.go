package services

import (
	"context"
	"gqlgen-go/graph/db"
	"gqlgen-go/graph/model"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type projectItemService struct {
	exec boil.ContextExecutor
}

func convertProjectV2Item(item *db.Projectcard) *model.ProjectV2Item {
	result := &model.ProjectV2Item{
		ID:      item.ID,
		Project: &model.ProjectV2{ID: item.Project},
	}
	if item.Issue.Valid {
		result.Content = &model.Issue{ID: item.Issue.String}
	}
	if item.Pullrequest.Valid {
		result.Content = &model.PullRequest{ID: item.Pullrequest.String}
	}
	return result
}

func (p *projectItemService) AddIssueInProjectV2(ctx context.Context, projectID, issueID string) (*model.ProjectV2Item, error) {
	itemID := uuid.New()
	item := &db.Projectcard{
		ID:      itemID.String(),
		Project: projectID,
		Issue:   null.StringFrom(issueID),
	}
	if err := item.Insert(ctx, p.exec, boil.Infer()); err != nil {
		return nil, err
	}
	return convertProjectV2Item(item), nil
}
