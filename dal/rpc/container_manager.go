package rpc

import (
	"context"
	"io"
	containerManager "web-IDE-back-end/proto/container_server"
	util "web-IDE-back-end/utils"

	"github.com/sirupsen/logrus"
)

// GetFile 获取container中的一个目录，tar格式
func GetFile(ctx context.Context, containerID, path string) (data []byte, err error) {
	logrus.Infof("container_id: %s", containerID)
	logrus.Infof("path: %s", path)
	req := &containerManager.GetFile_Request{ContainerId: containerID, Path: path}
	stream, err := containerManagerClient.GetFile(ctx, req)
	if err != nil {
		logrus.Warnf("GetFile failed, err: %+v", err)
		return
	}
	data = make([]byte, 0, 1<<10)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			logrus.Warnln(err)
			break
		}
		logrus.Infof("rpc GetFile get %d bytes", len(resp.Data))
		data = append(data, resp.Data...)
	}
	return
}

func GetDirectory(ctx context.Context, containerID, path string) (fileStat []*containerManager.FileStat, err error) {
	req := &containerManager.ListFile_Request{ContainerId: containerID, Path: path}
	resp, err := containerManagerClient.ListFile(ctx, req)
	if err != nil {
		logrus.Warnf("GetDirectory failed: %+v", err)
		return
	}
	return resp.Files, nil
}

func SaveFile(ctx context.Context, containerID, path, fileName, data string) (newVewsion string, err error) {
	// containerManager只接收tar文件
	tarFile, err := util.Pack([]util.File{{Name: fileName, Body: data}})

	if err != nil {
		return
	}
	req := &containerManager.UpdateFile_Request{
		ContainerId: containerID,
		Path:        path,
		FileName:    fileName,
		Force:       true,
		Data:        tarFile,
		OldVersion:  "", // todo 可以根据version检测文件冲突，后面再补
	}
	resp, err := containerManagerClient.UpdateFile(ctx, req)
	if err != nil {
		// logrus.Warnf("SaveFile failed: %+v", err)
		return
	}
	return resp.NewVersion, nil
}
