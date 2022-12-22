package authservice

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	sqlc "github.com/yards22/lcmanager/db/sqlc"
	kvstore "github.com/yards22/lcmanager/pkg/kv_store"
	util "github.com/yards22/lcmanager/util"
)

// session age in seconds
const SessionAge = 24 * 60 * 60

var ErrInternalServerError = errors.New("internal server error")
var ErrUnauthorized = errors.New("unauthorized")

type SendOTPArgs struct {
	MailId string `json:"mail_id"`
}

type LoginArgs struct {
	MailId string `json:"mail_id"`
	OTP    string `json:"otp"`
}

type AuthService struct {
	kv      *kvstore.RedisKVStore
	querier sqlc.Querier
}

func New(kv *kvstore.RedisKVStore, querier sqlc.Querier) *AuthService {
	return &AuthService{kv, querier}
}

func (as *AuthService) PerformMailIdCheck(ctx context.Context, arg SendOTPArgs) string {
	admin, err := as.querier.GetAdmin(ctx, arg.MailId)
	if err != nil {
		fmt.Println(err)
	}
	if len(admin) != 0 {
		otp := util.GenerateRandom(6)
		err := as.kv.Set("otp_"+otp, arg.MailId)
		fmt.Println(err)
		return otp
	}
	return uuid.Nil.String()
}

func (as *AuthService) PerformLogin(ctx context.Context, arg LoginArgs) string {
	userDetails := strings.Split(as.kv.Get("otp_"+arg.OTP), " ")[2]
	fmt.Println(userDetails)
	if userDetails == arg.MailId {
		admin, err := as.querier.GetAdmin(ctx, arg.MailId)
		if err != nil {
			fmt.Println(err)
		}
		categories := ""
		for i := 0; i < len(admin); i++ {
			categories += admin[i].OpenTo
			if i+1 != len(admin) {
				categories += "/"
			}
		}
		//    generate a token (n)
		token := util.GenerateRandom(16)
		//    set it into redis
		for as.kv.Get(token) == "Nil" {
			token = util.GenerateRandom(16)
		}
		as.kv.Set(token, categories)
		return token
	}
	return uuid.Nil.String()
}

func (as *AuthService) CheckSession(ctx context.Context, token string) []string {
	data := as.kv.Get(token)
	if data != uuid.Nil.String() {
		open_to := strings.Split(data, " ")[2]
		categories := strings.Split(open_to, "/")
		return categories
	}
	return nil
}
