package util

import (
	"archive/tar"
	"bytes"
	"io"
	"io/ioutil"
	"web-IDE-back-end/model"

	"github.com/sirupsen/logrus"
)

func Unpack(tarFile []byte) (files []*model.File, err error) {
	reader := bytes.NewReader(tarFile)
	tarReader := tar.NewReader(reader)
	for {
		hdr, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			logrus.Warnln(err)
			return files, err
		}
		logrus.Infof("unpack file: %s\n", hdr.Name)
		content, err := ioutil.ReadAll(tarReader)
		if err != nil {
			logrus.Warnln(err)
			return files, err
		}
		file := &model.File{Header: hdr, Content: string(content)}
		files = append(files, file)
	}
	return files, nil
}
