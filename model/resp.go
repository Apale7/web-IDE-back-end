package model

import (
	"archive/tar"
)

type File struct {
	Header  *tar.Header `json:"header"`
	Content string      `json:"content"`
}

type FileStat struct {
	FileType int32  `json:"file_type"`
	FileName string `json:"file_name"`
}
