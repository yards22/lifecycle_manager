package feedback_manager

import (
	"context"
	"fmt"

	sqlc "github.com/yards22/lcmanager/db/sqlc"
	"gopkg.in/guregu/null.v4"
)

type FeedbackManager struct {
	querier sqlc.Querier
}

type UpdateCommentsParams struct {
	Status     bool        `json:"status"`
	Comment    null.String `json:"comment"`
	FeedbackID int64       `json:"feedback_id"`
}

func New(querier sqlc.Querier) *FeedbackManager {
	return &FeedbackManager{querier}
}

func (fm *FeedbackManager) GetFeedback(ctx context.Context) []*sqlc.Feedback {

	feedback, err := fm.querier.GetFeedback(ctx)

	if err != nil {
		fmt.Println(err)
	}

	return feedback

}

func (fm *FeedbackManager) PostCommentFeedback(ctx context.Context, arg UpdateCommentsParams) {

	err := fm.querier.UpdateComments(ctx, sqlc.UpdateCommentsParams{
		Status:     arg.Status,
		Comment:    arg.Comment,
		FeedbackID: arg.FeedbackID,
	})

	if err != nil {
		fmt.Println(err)
	}

	return

}
