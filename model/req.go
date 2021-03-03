package model

// JudgeReq 用户提交代码时的请求体
type JudgeReq struct {
	Base
	Language string `json:"language"`
	Input    string `json:"input"`
	Code     string `json:"code"`
}

type LoginReq struct {
	Base
	Username string `json:"username"`
	Password string `json:"password"`
}
