package handler

import (
	"web-IDE-back-end/model"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Judge(c *gin.Context) {
	var req model.JudgeReq
	if err := c.Bind(&req); err != nil {
		logrus.Warnf("Invalid params: %v", err)
		c.String(200, "Invalid params")
		return
	}
	logrus.Infof("Get JudgeReq: %+v", req)
	c.String(200, "OK")
}
