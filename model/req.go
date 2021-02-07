package model

// JudgeReq 用户提交代码时的请求体
type JudgeReq struct {
	Language string `json:"language"`
	Input    string `json:"input"`
	Code     string `json:"code"`
}
