package handler

import (
	"context"
	"errors"
	"fmt"
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
	if err := c.ShouldBind(&reqBody); err != nil {
		logrus.Warnf("invalid params, err: %v", err)
		utils.RetErr(c, constdef.ErrInvalidParams)
		return
	}
	logrus.Infof("Get reqBody: %+v", reqBody)
	tarFile, err := rpc.GetFile(ctx, reqBody.ContainerID, reqBody.Path)
	if err != nil {
		logrus.Warnf("GetFile failed, err: %v", err)
		utils.RetErr(c, errors.New("GetFile error"))
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

//GetDir 获取一个目录的文件信息(不包括内容)
func GetDir(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.GetFileReq
	if err := c.ShouldBind(&reqBody); err != nil {
		logrus.Warnf("invalid params, err: %v", err)
		utils.RetErr(c, constdef.ErrInvalidParams)
		return
	}
	logrus.Infof("Get reqBody: %+v", reqBody)

	fileStat, err := rpc.GetDirectory(ctx, reqBody.ContainerID, reqBody.Path)
	if err != nil {
		logrus.Warnf("GetDirectory failed, err: %v", err)
		utils.RetErr(c, errors.New("GetDirectory error"))
		return
	}
	fmt.Printf("Get %d files\n", len(fileStat))
	resFileStat := make([]*model.FileStat, 0, len(fileStat))
	for _, f := range fileStat {
		resFileStat = append(resFileStat, f.ToRespFileStat())
	}
	utils.RetData(c, resFileStat)
}
