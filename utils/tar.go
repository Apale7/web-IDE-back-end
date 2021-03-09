package util

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
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

type File struct {
	Name, Body string
}

func Pack(files []File) (res []byte, err error) {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			logrus.Fatal(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			logrus.Fatal(err)
		}
	}
	if err := tw.Close(); err != nil {
		logrus.Fatal(err)
	}

	// //查看压缩内容
	// f, err := os.Create("temp_dir/a.tar")
	// if err != nil {
	// 	return
	// }
	// n, err := buf.WriteTo(f)
	// if err != nil {
	// 	return
	// }
	// fmt.Println("size:", n)
	res = buf.Bytes()
	tr := tar.NewReader(&buf)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // 文件已经遍历完成了
		}
		if err != nil {
			logrus.Fatal(err)
		}
		fmt.Printf("%s的文件内容: ", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			logrus.Fatal(err)
		}
		fmt.Println()
	}
	return
}
