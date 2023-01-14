package t_users_manager

import (
	"context"
	"fmt"
	"log"
	"sort"
	"time"

	sqlc "github.com/yards22/lcmanager/db/sqlc"
	"github.com/yards22/lcmanager/pkg/runner"
)

type TUManager struct {
	querier sqlc.Querier
	runner  *runner.Runner
}

type Tentries struct {
	id    int32
	count int
}

func New(querier sqlc.Querier, interval time.Duration) *TUManager {
	log.Println("setup trending user runner at interval", interval.Minutes())
	return &TUManager{querier, runner.New(interval)}
}

// Function to generate trending users .
func (tpm *TUManager) GenerateTrendingUsers(ctx context.Context) {

	// Get Likes-Posts array.
	// These are all the likes that got generated during last day.

	likes, err := tpm.querier.LikeTrendingUsers(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Get Comment-Posts array.
	// These are all the likes that got generated during last day.

	comments, err := tpm.querier.CommentTrendingUsers(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Get Share-Posts array.
	// Will be implemented after the share functionality is added.

	// Combine all the above and form CAI and post array.

	CAI := make(map[int32]int)

	for _, j := range likes {
		CAI[(*j).UserID] += int((*j).LikeCount)
	}

	for _, j := range comments {
		CAI[(*j).UserID] += int((*j).CommentCount)
	}

	keys := make([]int32, 0, len(CAI))
	for key := range CAI {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return CAI[keys[i]] > CAI[keys[j]]
	})

	// pick top 10 of the above array.

	topPicks := make([]Tentries, 0, 25)

	for _, k := range keys {
		topPicks = append(topPicks, Tentries{k, CAI[k]})
	}

	fmt.Println(topPicks)

	// Insert these posts into Trending Table.
	for _, j := range topPicks {
		tpm.querier.InsertTrendingUsers(ctx, j.id)
	}

}

func (tm *TUManager) Run() {
	tm.runner.Run(func() {
		log.Println("invoking trending user runner fn")
		tm.GenerateTrendingUsers(context.Background())
	})

}

func (tm *TUManager) Close() {
	tm.runner.Close()
}
