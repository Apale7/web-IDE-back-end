package model

import (
	"archive/tar"
	user_center "web-IDE-back-end/proto/user-center"
)

type File struct {
	Header  *tar.Header `json:"header"`
	Content string      `json:"content"`
}

type FileStat struct {
	FileType int32  `json:"file_type"`
	FileName string `json:"file_name"`
}

type LoginResp struct {
	user_center.LoginResponse
	Auth []string `json:"auth"`
}
