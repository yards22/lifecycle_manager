package poll_manager

import (
	"context"
	"fmt"

	sqlc "github.com/yards22/lcmanager/db/sqlc"
)

type PollManager struct {
	querier sqlc.Querier
}

func New(querier sqlc.Querier) *PollManager {
	return &PollManager{querier}
}

func (pm *PollManager) Create(ctx context.Context, arg sqlc.CreatePollsParams) {
	err := pm.querier.CreatePolls(ctx, sqlc.CreatePollsParams{
		PollQuestion: arg.PollQuestion,
		OptionsCount: arg.OptionsCount,
		Options:      arg.Options,
	})

	if err != nil {
		fmt.Println(err)
	}

}

func (pm *PollManager) Get(ctx context.Context) []*sqlc.Poll {
	polls, err := pm.querier.GetPolls(ctx)

	if err != nil {
		fmt.Println(err)
	}

	return polls

}
