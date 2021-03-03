package handler

import (
	"context"
	"errors"
	"strconv"
	"web-IDE-back-end/dal/rpc"
	"web-IDE-back-end/model"
	"web-IDE-back-end/proto/codeRunner"

	"github.com/Apale7/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//Judge for Community Edition
func Judge(c *gin.Context) {
	ctx := context.Background()
	var req model.JudgeReq
	if err := c.ShouldBind(&req); err != nil {
		logrus.Warnf("Invalid params: %v", err)
		utils.RetErr(c, errors.New("Invalid params"))
		return
	}
	logrus.Infof("Get JudgeReq: %+v", req)
	language, _ := strconv.Atoi(req.Language)
	resp, err := rpc.Judge(ctx, CodeRunner.CodeLanguage(language), req.Input, req.Code)
	if err != nil {
		logrus.Warnf("server error: %v", err)
		utils.RetErr(c, errors.New("server error"))
		return
	}
	c.JSON(200, gin.H{"code": "0", "resp": resp})
}
