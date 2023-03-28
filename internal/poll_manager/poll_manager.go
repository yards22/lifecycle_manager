package poll_manager

import (
	"context"
	"fmt"
	
	sqlc "github.com/yards22/lcmanager/db/sqlc"
)

type PollManager struct {
	querier sqlc.Querier
}

type PollsRefinedFormat struct{
     
}
func New(querier sqlc.Querier) *PollManager {
	return &PollManager{querier}
}

func (pm *PollManager) Create(ctx context.Context, arg sqlc.CreatePollsParams) {
	err := pm.querier.CreatePolls(ctx, sqlc.CreatePollsParams{
		PollBy:       arg.PollBy,
		PollQuestion: arg.PollQuestion,
		OptionsCount: arg.OptionsCount,
		Options:      arg.Options,
	})

	if err != nil {
		fmt.Println(err)
	}

}


func (pm *PollManager) Get(ctx context.Context) []*sqlc.GetPollsRow {
	polls, err := pm.querier.GetPolls(ctx)

	if err != nil {
		fmt.Println(err)
	}

	for i:=0;i<len(polls);i++ {
        
	}

	return polls

}
