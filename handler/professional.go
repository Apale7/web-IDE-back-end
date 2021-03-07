package handler

import (
	"context"
	"errors"
	"web-IDE-back-end/dal/rpc"
	"web-IDE-back-end/model"
	util "web-IDE-back-end/utils"

	"github.com/Apale7/common/constdef"
	"github.com/Apale7/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//GetFile 获取一个文件的内容
func GetFile(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.GetFileReq
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		logrus.Warnf("invalid params, err: %v", err)
		utils.RetErr(c, constdef.ErrInvalidParams)
		return
	}
	logrus.Infof("Get reqBody: %+v", reqBody)
	tarFile, err := rpc.GetFile(ctx, reqBody.ContainerID, reqBody.Path)
	if err != nil {
		return
	}
	files, err := util.Unpack(tarFile)
	if err != nil {
		logrus.Warnf("Unpack failed, err: %v", err)
		utils.RetErr(c, errors.New("Unpack error"))
		return
	}
	utils.RetData(c, gin.H{"files": files})
}

//GetDir 获取一个目录
func GetDir(c *gin.Context) {
	var reqBody model.GetFileReq
	if err := c.ShouldBind(&reqBody); err != nil {
		logrus.Warnf("invalid params, err: %v", err)
		utils.RetErr(c, constdef.ErrInvalidParams)
		return
	}
}
