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

type RefreshReq struct {
	Base
	RefreshToken string `json:"refresh_token"`
}

// GetFileReq GetFile和GetDir公用
type GetFileReq struct {
	ContainerID string `json:"container_id" form:"container_id"`
	Path        string `json:"path" form:"path"`
}
