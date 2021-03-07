package rpc

import (
	"archive/tar"
	"context"
	"fmt"
	"testing"
	util "web-IDE-back-end/utils"

	"github.com/sirupsen/logrus"
)

func TestGetDirectory(t *testing.T) {

}

var ctx = context.Background()

func TestGetFile(t *testing.T) {
	tarFile, err := GetFile(ctx, "container2", "/etc/apt")
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
