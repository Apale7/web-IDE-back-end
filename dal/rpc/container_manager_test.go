package rpc

import (
	"archive/tar"
	"context"
	"fmt"
	"testing"
	util "web-IDE-back-end/utils"

	"github.com/sirupsen/logrus"
)
var ctx = context.Background()

// func TestGetDirectory(t *testing.T) {
// 	GetDirectory(ctx, "")
// }


func TestGetFile(t *testing.T) {
	tarFile, err := GetFile(ctx, "container3", "/root/a.py")
	if err != nil {
		logrus.Warnln(err)
		t.FailNow()
	}
	files, err := util.Unpack(tarFile)
	for _, f := range files {
		fmt.Println(f.Header.Name, f.Header.Typeflag)
		if f.Header.Typeflag == tar.TypeReg {
			fmt.Println(f.Content)
		}
	}
}
