package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	// if os.Getenv("ENV") != "dev" {
	// 	logrus.SetLevel(logrus.WarnLevel)
	// }
}

func main() {

	r := gin.Default()
	defer r.Run(":3456")
	collectRoutes(r)
}
