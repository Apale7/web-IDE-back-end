package handler

import (
	"context"
	"web-IDE-back-end/dal/rpc"
	"web-IDE-back-end/model"

	"github.com/Apale7/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Login(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.LoginReq
	if err := c.ShouldBind(&reqBody); err != nil {
		logrus.Warnf("invalid params: %v", err)
		utils.RetErr(c, err)
		return
	}
	logrus.Infof("reqBody: %+v", reqBody)
	accessToken, refreshToken, err := rpc.Login(ctx, reqBody.Username, reqBody.Password)
	if err != nil {
		logrus.Warnf("login failed, err: %v", err)
		utils.RetErr(c, err)
		return
	}
	utils.RetData(c, gin.H{"accessToken": accessToken, "refreshToken": refreshToken})
}

func Register(c *gin.Context) {

}

func Refresh(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.RefreshReq
	if err := c.ShouldBind(&reqBody); err != nil {
		logrus.Warnf("invalid params: %v", err)
		utils.RetErr(c, err)
		return
	}
	logrus.Infof("reqBody: %+v", reqBody)

	accessToken, refreshToken, err := rpc.Refresh(ctx, reqBody.RefreshToken)
	if err != nil {
		logrus.Warnf("login failed, err: %v", err)
		utils.RetErr(c, err)
		return
	}
	utils.RetData(c, gin.H{"accessToken": accessToken, "refreshToken": refreshToken})
}
