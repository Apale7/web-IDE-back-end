package rpc

import (
	"context"
	"web-IDE-back-end/proto/CodeRunner"

	"github.com/sirupsen/logrus"
)

// Judge Judge the code and return the result
func Judge(ctx context.Context, language CodeRunner.CodeLanguage, input, code string) (resp *CodeRunner.CodeRunnerRespone, err error) {
	req := &CodeRunner.CodeRunnerRequest{
		Language_: language,
		Input_:    input,
		Code_:     code,
	}
	resp, err = codeRunnerClient.Judge(ctx, req)
	if err != nil {
		logrus.Warnf("Judge error: %+v", err)
		return
	}
	return
}
