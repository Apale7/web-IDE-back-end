package handler

import (
	"context"
	"web-IDE-back-end/dal/rpc"
	"web-IDE-back-end/model"
	"web-IDE-back-end/proto/CodeRunner"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Judge(c *gin.Context) {
	ctx := context.Background()
	var req model.JudgeReq
	if err := c.Bind(&req); err != nil {
		logrus.Warnf("Invalid params: %v", err)
		c.String(200, "Invalid params")
		return
	}
	logrus.Infof("Get JudgeReq: %+v", req)
	resp, _ := rpc.Judge(ctx, CodeRunner.CodeLanguage(req.Language), req.Input, req.Code)
	// if err != nil {
	// 	logrus.Warnf("server error: %v", err)
	// 	c.String(200, "服务器错误")
	// 	return
	// }
	c.JSON(200, gin.H{"code": "0", "resp": resp})
}
