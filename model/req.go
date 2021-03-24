package model

// JudgeReq 用户提交代码时的请求体
type JudgeReq struct {
	Base
	Language string `json:"language" form:"language"`
	Input    string `json:"input" form:"input"`
	Code     string `json:"code" form:"code"`
}

type LoginReq struct {
	Base
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type RefreshReq struct {
	Base
	RefreshToken string `json:"refresh_token" form:"refresh_token"`
}

// GetFileReq GetFile和GetDir公用
type GetFileReq struct {
	ContainerID string `json:"container_id" form:"container_id"`
	Path        string `json:"path" form:"path"`
}

type SaveFileReq struct {
	ContainerID string `json:"container_id" form:"container_id"`
	Path        string `json:"path" form:"path"`
	Data        string `json:"data" form:"data"`
}
