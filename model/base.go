package model

type Base struct {
	//签名
	Sign string `json:"sign" form:"sign"`
	//发起请求的时间戳
	ReqTime string `json:"reqTime" form:"reqTime"`
}
