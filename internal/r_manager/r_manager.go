package r_manager

import (
	"context"
	"fmt"
	"log"
	"math"
	"time"

	sqlc "github.com/yards22/lcmanager/db/sqlc"
	"github.com/yards22/lcmanager/pkg/app_config"
	"github.com/yards22/lcmanager/pkg/runner"
)

type Categories struct {
	posts             float64
	started_following float64
	added_followers   float64
	reactions         float64
}

type RatingManager struct {
	querier sqlc.Querier
	runner  *runner.Runner
}

func New(querier sqlc.Querier, interval time.Duration) *RatingManager {
	log.Println("setup rating runner at interval", interval.Minutes())
	return &RatingManager{querier, runner.New(interval)}
}

func (rm *RatingManager) UpdateRatings(ctx context.Context) {
	// create a map for storing this week credits of each user.
	score := make(map[int]float64)

	usersCount, err := rm.querier.GetUsers(ctx)

	fmt.Println(usersCount)

	proportions := Categories{
		posts:             35,
		started_following: 5,
		added_followers:   30,
		reactions:         30,
	}

	idealScores := Categories{
		posts:             7,
		started_following: 10,
		added_followers:   15,
		reactions:         2,
	}

	if err != nil {
		fmt.Println(err)
	}

	duration_ := app_config.Data.MustInt("duration_rating")

	// score changes because of Posts...

	posts, err := rm.querier.GetPosts(ctx, duration_)

	if err != nil {
		fmt.Println(err)
	}

	for post := 0; post < len(posts); post++ {
		p := proportions.posts
		i_s := idealScores.posts
		res_posts := math.Min(float64(posts[post].Count), i_s)
		score[int(posts[post].UserID)] += res_posts * p / i_s
	}

	// // score changes because of followers & following .

	followers, err := rm.querier.GetFollwers(ctx, duration_)

	if err != nil {
		fmt.Println(err)
	}

	for follower := 0; follower < len(followers); follower++ {
		p := proportions.added_followers
		i_s := idealScores.added_followers
		res_followers := math.Min(float64(followers[follower].Count), i_s)
		score[int(followers[follower].UserID)] += res_followers * p / i_s
	}

	following, err := rm.querier.GetFollowing(ctx, duration_)

	if err != nil {
		fmt.Println(err)
	}

	for following_ := 0; following_ < len(following); following_++ {
		p := proportions.started_following
		i_s := idealScores.started_following
		res_following := math.Min(float64(followers[following_].Count), i_s)
		score[int(following[following_].UserID)] += res_following * p / i_s
	}

	// score changes because of reactions

	p := proportions.reactions
	i_s := idealScores.reactions

	likes, err := rm.querier.GetUserLikes(ctx, duration_)

	if err != nil {
		fmt.Println(err)
	}

	comments, err := rm.querier.GetUserComments(ctx, duration_)

	if err != nil {
		fmt.Println(err)
	}

	reactions := make(map[int]float64)

	for like_ := 0; like_ < len(likes); like_++ {
		fmt.Println(likes[like_].UserID, likes[like_].LikeCount)
		reactions[int(likes[like_].UserID)] += float64(likes[like_].LikeCount)
	}

	for comment_ := 0; comment_ < len(comments); comment_++ {
		reactions[int(comments[comment_].UserID)] += float64(comments[comment_].CommentCount)
	}

	// get followers count

	follower_count, err := rm.querier.GetFollowersCount(ctx)

	for key, value := range reactions {

		if follower_count[key-1] != 0 {
			temp := value / float64(follower_count[key-1])
			res_reaction := math.Min(temp, i_s)
			score[key] += res_reaction * p / i_s
		} else {
			score[key] += p
		}
	}

	if err != nil {
		fmt.Println(err)
	}

	// Rating Updation ....

	for user := 1; user <= int(usersCount); user++ {

		get_rating, err := rm.querier.GetRating(ctx, int32(user))

		if err != nil {
			println(err)
		}

		updated_index := get_rating + int32(rm.RatingFunction(score[user], (get_rating/200)))

		rm.querier.UpdateRating(ctx, sqlc.UpdateRatingParams{
			CricIndex: updated_index,
			UserID:    int32(user),
		})
	}

}

func (rm *RatingManager) RatingFunction(score float64, present_slab int32) float64 {

	threshold := [9]float64{0, 5, 10, 15, 20, 25, 30, 35}

	denom := math.Log2(float64(present_slab + 2))
	num := score - threshold[present_slab]
	return num / denom

}

func (rm *RatingManager) Run() {
	rm.runner.Run(func() {
		log.Println("invoking rating runner fn")
		rm.UpdateRatings(context.Background())
	})
}

func (rm *RatingManager) Close() {
	rm.runner.Close()
}
