package main

import (
	"context"
	"web-IDE-back-end/dal/rpc"
	user_center "web-IDE-back-end/proto/user-center"
)

func main() {
	ctx := context.Background()
	user := user_center.User{
		Username: "apale",
		Password: "123465",
	}
	extra := user_center.UserExtra{
		Email:       "1092377056@qq.com",
		Nickname:    "Apale",
		PhoneNumber: "13710247812",
	}
	rpc.Register(ctx, user, extra)
}
