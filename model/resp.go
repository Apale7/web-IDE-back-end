package model

import "archive/tar"

type File struct {
	Header  *tar.Header `json:"header"`
	Content string      `json:"content"`
}
