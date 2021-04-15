package rpc

import (
	codeRunner "web-IDE-back-end/proto/codeRunner"
	containerManager "web-IDE-back-end/proto/container_server"
	user_center "web-IDE-back-end/proto/user-center"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	userCenterClient       user_center.UserCenterClient
	codeRunnerClient       codeRunner.CodeRunnerClient
	containerManagerClient containerManager.ManagerClient
)

func init() {
	userCenterClient = user_center.NewUserCenterClient(getConn("111.230.172.240:9999"))
	// userCenterClient = user_center.NewUserCenterClient(getConn(":9999"))
	codeRunnerClient = codeRunner.NewCodeRunnerClient(getConn("193.112.177.167:8233"))
	containerManagerClient = containerManager.NewManagerClient(getConn(":8666"))
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
