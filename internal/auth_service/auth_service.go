package authservice

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	sqlc "github.com/yards22/lcmanager/db/sqlc"
	"github.com/yards22/lcmanager/pkg/app_config"
	kvstore "github.com/yards22/lcmanager/pkg/kv_store"
	"github.com/yards22/lcmanager/pkg/mailer"
	util "github.com/yards22/lcmanager/util"
)

// session age in seconds
const SessionAge = 24 * 60 * 60

var ErrInternalServerError = errors.New("internal server error")
var ErrUnauthorized = errors.New("unauthorized")

type BlogsCxt struct {
	Blogs bool
}
type PollsCxt struct {
	Polls bool
}
type SendOTPArgs struct {
	MailId string `json:"mail_id"`
}

type LoginArgs struct {
	MailId string `json:"mail_id"`
	OTP    string `json:"otp"`
}
type InsertAdminParams struct {
	MailID string   `json:"mail_id"`
	OpenTo []string `json:"open_to"`
}

type RegisterRoleArgs struct {
	MailId    string `json:"mail_id"`
	Role      string `json:"role"`
	SecretKey string `json:"secret_key"`
}

type AuthService struct {
	kv      *kvstore.RedisKVStore
	querier sqlc.Querier
	mailer  *mailer.GoMail
}

func New(kv *kvstore.RedisKVStore, querier sqlc.Querier, mailer *mailer.GoMail) *AuthService {
	return &AuthService{kv, querier, mailer}
}

func (as *AuthService) PerformMailIdCheck(ctx context.Context, arg SendOTPArgs) string {
	admin, err := as.querier.GetAdmin(ctx, arg.MailId)
	if err != nil {
		fmt.Println(err)
	}
	if len(admin) != 0 {
		otp := util.GenerateRandom(6)
		err := as.kv.Set("admin_otp_"+otp, arg.MailId)
		fmt.Println(err)
		err = as.mailer.Send("22Yardz Admin", arg.MailId, "OTP Verification", fmt.Sprintf("Your OTP is %s.", otp))
		if err != nil {
			fmt.Println(err)
		}
		return otp
	}
	return uuid.Nil.String()
}

func (as *AuthService) PerformLogin(ctx context.Context, arg LoginArgs) string {
	userDetails := strings.Split(as.kv.Get("admin_otp_"+arg.OTP), " ")[2]
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
		token := util.GenerateRandomToken(64)
		for as.kv.Get(token) == "Nil" {
			token = util.GenerateRandomToken(64)
		}
		x := arg.MailId + " " + categories
		as.kv.Set("admin_"+token, x)
		return token
	}
	return uuid.Nil.String()
}

func (as *AuthService) PerformRoleSignup(ctx context.Context, arg RegisterRoleArgs) string {
	secret_key := app_config.Data.MustString("SECRET")
	if arg.SecretKey == secret_key {
		err := as.querier.InsertAdmin(ctx, sqlc.InsertAdminParams{
			MailID: arg.MailId,
			OpenTo: arg.Role,
		})
		if err != nil {
			fmt.Println(err)
		}
		return "successful_signup"
	}
	return uuid.Nil.String()
}

func (as *AuthService) PerformLogout(ctx context.Context, token string) {
	as.kv.Delete("admin_" + token)
}

func (as *AuthService) CheckSession(ctx context.Context, token string) InsertAdminParams {
	data := as.kv.Get("admin_" + token)
	fmt.Println("checkSession data ", data)
	var params InsertAdminParams
	if data != uuid.Nil.String() {
		open_to := strings.Split(data, " ")[3]
		params.MailID = strings.Split(data, " ")[2]
		params.OpenTo = strings.Split(open_to, "/")
	}
	return params
}
