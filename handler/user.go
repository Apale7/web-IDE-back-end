package handler

import (
	"context"
	"web-IDE-back-end/dal/rpc"
	"web-IDE-back-end/model"
	user_center "web-IDE-back-end/proto/user-center"

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
	resp, err := rpc.Login(ctx, reqBody.Username, reqBody.Password)
	if err != nil {
		logrus.Warnf("login failed, err: %v", err)
		utils.RetErr(c, err)
		return
	}
	var loginRes model.LoginResp
	loginRes.LoginResponse = *resp
	loginRes.Auth = []string{"super"}
	utils.RetData(c, loginRes)
}

func Register(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.RegisterReq
	if err := c.ShouldBind(&reqBody); err != nil {
		logrus.Warnf("invalid params: %v", err)
		utils.RetErr(c, err)
		return
	}
	logrus.Infof("reqBody: %+v", reqBody)

	err := rpc.Register(ctx, &user_center.User{Username: reqBody.Username, Password: reqBody.Password}, &user_center.UserExtra{Nickname: reqBody.Nickname, PhoneNumber: reqBody.PhoneNumber})

	if err != nil {
		logrus.Warnf("register failed, err: %v", err)
		utils.RetErr(c, err)
		return
	}
	utils.RetSuccess(c)
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

	resp, err := rpc.Refresh(ctx, reqBody.RefreshToken)
	if err != nil {
		logrus.Warnf("refresh failed, err: %v", err)
		utils.RetErr(c, err)
		return
	}
	utils.RetData(c, resp)
}