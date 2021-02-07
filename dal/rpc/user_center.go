package rpc

import (
	"context"
	user_center "web-IDE-back-end/proto/user-center"

	"github.com/sirupsen/logrus"
)

func Register(ctx context.Context, user user_center.User, extra user_center.UserExtra) (err error) {
	req := &user_center.RegisterRequest{
		User:      &user,
		UserExtra: &extra,
	}
	resp, err := userCenterClient.Register(ctx, req)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	logrus.Infof("注册成功, id: %d\n", resp.Id)
	return
}
