package rpc

import (
	"context"
	user_center "web-IDE-back-end/proto/user-center"

	"github.com/sirupsen/logrus"
)

// Register 注册
func Register(ctx context.Context, user *user_center.User, extra *user_center.UserExtra) (err error) {
	req := &user_center.RegisterRequest{
		User:      user,
		UserExtra: extra,
	}
	resp, err := userCenterClient.Register(ctx, req)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	logrus.Infof("注册成功, id: %d\n", resp.Id)
	return
}

func Login(ctx context.Context, username, password string) (token string, err error) {
	req := &user_center.LoginRequest{
		Username: username,
		Password: password,
	}
	resp, err := userCenterClient.Login(ctx, req)
	if err != nil {
		return
	}
	return resp.Token, nil
}

func CheckToken(ctx context.Context, token string) (userInfo *user_center.UserInfo, err error) {
	req := &user_center.CheckTokenRequest{Token: token}
	resp, err := userCenterClient.CheckToken(ctx, req)
	if err != nil {
		return
	}
	return resp.User, nil
}
