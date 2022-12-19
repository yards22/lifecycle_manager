package feedback_manager

import (
	"context"
	"fmt"

	sqlc "github.com/yards22/lcmanager/db/sqlc"
)

type FeedbackManager struct {
	querier sqlc.Querier
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
