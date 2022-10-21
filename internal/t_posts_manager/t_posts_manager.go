package t_posts_manager

import (
	"context"
	"fmt"
	"sort"
	"time"

	sqlc "github.com/yards22/lcmanager/db/sqlc"
	"github.com/yards22/lcmanager/pkg/runner"
)

type TPManager struct {
	querier sqlc.Querier
	runner  *runner.Runner
}

type Tentries struct {
	id    int64
	count int
}

func New(querier sqlc.Querier, interval time.Duration) *TPManager {
	return &TPManager{querier, runner.New(interval)}
}

// Function to clean expired tokens from DB
func (tpm *TPManager) GenerateTrendingPosts(ctx context.Context) {

	// Get Likes-Posts array.
	// These are all the likes that got generated during last week.

	likes, err := tpm.querier.LikeTrending(ctx)
	if err != nil {
		fmt.Println(err.Error())

	}

	// Get Comment-Posts array.
	// These are all the likes that got generated during last week.

	comments, err := tpm.querier.CommentTrending(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Get Share-Posts array.

	// Combine all the above and form CAI and post array.

	CAI := make(map[int64]int)

	for _, j := range likes {
		CAI[(*j).PostID] += int((*j).LikeCount)

	}

	for _, j := range comments {
		CAI[(*j).PostID] += int((*j).CommentCount)
	}

	keys := make([]int64, 0, len(CAI))
	for key := range CAI {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return CAI[keys[i]] > CAI[keys[j]]
	})

	// pick top 10 of the above array.

	topPicks := make([]Tentries, 0, 10)

	for _, k := range keys {
		topPicks = append(topPicks, Tentries{k, CAI[k]})
	}

	fmt.Println(topPicks)

	// Insert these posts into Trending Table.
	for _, j := range topPicks {
		tpm.querier.InsertTrending(ctx, j.id)
	}

}

func (tm *TPManager) Run() {
	tm.runner.Run(func() {
		tm.GenerateTrendingPosts(context.Background())
	})

}

func (tm *TPManager) Close() {
	tm.runner.Close()
}
