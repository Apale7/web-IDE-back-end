package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	port string
)

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	port = os.Getenv("web_ide_back_end_port")
	if port == "" {
		port = "3456"
	}
	// if os.Getenv("ENV") != "dev" {
	// 	logrus.SetLevel(logrus.WarnLevel)
	// }
}

func main() {
	r := gin.Default()
	defer r.Run(":" + port)
	collectRoutes(r)
}
