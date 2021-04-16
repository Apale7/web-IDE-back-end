package rpc

import (
	"fmt"
	"os"
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

var (
	userCenterAddr       string
	codeRunnerAddr       string = "[::]:8233"
	containerManagerAddr string = "[::]:8666"
)

func init() {
	initAddrs() //初始化各服务的地址
	userCenterClient = user_center.NewUserCenterClient(getConn(userCenterAddr))
	// userCenterClient = user_center.NewUserCenterClient(getConn(":9999"))
	codeRunnerClient = codeRunner.NewCodeRunnerClient(getConn(codeRunnerAddr))
	containerManagerClient = containerManager.NewManagerClient(getConn(containerManagerAddr))
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

func initAddrs() {
	userCenterAddr = os.Getenv("user_center_addr")
	if userCenterAddr == "" {
		userCenterAddr = ":9999"
	}
	fmt.Println(userCenterAddr)
	fmt.Println(containerManagerAddr)
	fmt.Println(codeRunnerAddr)
}
