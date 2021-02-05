package rpc

import (
	"context"
	user_center "web-IDE-back-end/proto/user-center"

	"github.com/sirupsen/logrus"
)

func Register(ctx context.Context, req *user_center.RegisterRequest) (resp *user_center.RegisterResponse, err error){
	resp, err = UserCenterClient.Register(ctx, req)
	if err != nil {
		logrus.Warnf("Register error: %+v", err)
	}
	return resp, err
}
