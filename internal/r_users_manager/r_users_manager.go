package r_users_manager

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	sqlc "github.com/yards22/lcmanager/db/sqlc"
	"github.com/yards22/lcmanager/pkg/runner"
)

type RUManager struct {
	querier sqlc.Querier
	runner  *runner.Runner
}
type UpsertRecommendationsParams struct {
	UserID      int32  `json:"user_id"`
	Recommend   string `json:"recommend"`
	Recommend_2 string `json:"recommend_2"`
}

func New(querier sqlc.Querier, interval time.Duration) *RUManager {
	return &RUManager{querier, runner.New(interval)}
}

// Function to Upsertrecommended users
func (rum *RUManager) GenerateRecommendedUsers(ctx context.Context) {

	usersCount, err := rum.querier.GetUsers(ctx)
	if err != nil {
		fmt.Println(err)
	}

	for user := 1; user <= int(usersCount); user++ {

		mutuals_id, err := rum.querier.GetMutual(ctx, int32(user))

		if err != nil {
			fmt.Println(err)
		}

		var stringified_mutual_id []string

		for mutual := 0; mutual < len(mutuals_id); mutual++ {
			m_id := strconv.Itoa(int(mutuals_id[mutual]))
			stringified_mutual_id = append(stringified_mutual_id, m_id)
		}

		res := strings.Join(stringified_mutual_id, "-")

		if res != "" {
			rum.querier.UpsertUserRecommendations(ctx, sqlc.UpsertUserRecommendationsParams{
				UserID:      int32(user),
				Recommend:   res,
				Recommend_2: res,
			})
		}

	}

}

func (rum *RUManager) Run() {
	rum.runner.Run(func() {
		rum.GenerateRecommendedUsers(context.Background())
	})

}

func (rum *RUManager) Close() {
	rum.runner.Close()
}
