package Manage

import "web-IDE-back-end/model"

func (f *FileStat) ToRespFileStat() *model.FileStat {
	return &model.FileStat{
		FileType: int32(f.FileType),
		FileName: f.FileName,
	}
}
