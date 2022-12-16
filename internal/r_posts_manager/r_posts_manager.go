package r_posts_manager

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	sqlc "github.com/yards22/lcmanager/db/sqlc"
	"github.com/yards22/lcmanager/pkg/runner"
)

type RPManager struct {
	querier sqlc.Querier
	runner  *runner.Runner
}

func New(querier sqlc.Querier, interval time.Duration) *RPManager {
	return &RPManager{querier, runner.New(interval)}
}

// Function to Generate Recommended Posts from DB
func (rpm *RPManager) GenerateRecommendedPosts(ctx context.Context) {
	// get the posts on which users following has reacted.
	usersCount, err := rpm.querier.GetUsers(ctx)

	if err != nil {
		fmt.Println(err)
	}

	for user := 1; user <= int(usersCount); user++ {

		post_id, err := rpm.querier.GetFollowingReaction(ctx, int32(user))

		var stringified_post_id []string

		for post := 0; post < len(post_id); post++ {
			p_id := strconv.Itoa(int(post_id[post]))
			stringified_post_id = append(stringified_post_id, p_id)
		}

		if err != nil {
			fmt.Println(err)
		}

		// stringify the list of post_id's

		res := strings.Join(stringified_post_id, "-")
		if res != "" {
			rpm.querier.UpsertPostRecommendations(ctx, sqlc.UpsertPostRecommendationsParams{
				UserID:                int32(user),
				PostRecommendations:   res,
				PostRecommendations_2: res,
			})
		}

	}

}

func (rm *RPManager) Run() {
	rm.runner.Run(func() {
		rm.GenerateRecommendedPosts(context.Background())
	})

}

func (rm *RPManager) Close() {
	rm.runner.Close()
}
