package rpc

import (
	"web-IDE-back-end/proto/user-center"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)
var (
	UserCenterClient user_center.UserCenterClient
)

func init() {
	UserCenterClient = user_center.NewUserCenterClient(getConn("localhost:8888"))
}

func getConn(addr string) *grpc.ClientConn{
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("fail to dial: %+v", err)
	}
	return conn
}
