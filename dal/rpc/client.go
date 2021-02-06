package rpc

import (
	"web-IDE-back-end/proto/CodeRunner"
	user_center "web-IDE-back-end/proto/user-center"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	userCenterClient user_center.UserCenterClient
	codeRunnerClient CodeRunner.CodeRunnerClient
)

func init() {
	userCenterClient = user_center.NewUserCenterClient(getConn("localhost:8888"))
	codeRunnerClient = CodeRunner.NewCodeRunnerClient(getConn("114.215.196.246:8233"))
}

func getConn(addr string) *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		logrus.Fatalf("%+v", err)
	}
	return conn
}
