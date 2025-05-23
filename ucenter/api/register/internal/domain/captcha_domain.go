package domain

import (
	"encoding/json"

	"github.com/wushengyouya/coin_exchange/coin-common/tools"
	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaDomain struct{}

type vaptchaReq struct {
	Id        string `json:"id"`
	SecretKey string `json:"secretkey"`
	Scene     int    `json:"scene"`
	Token     string `json:"token"`
	Ip        string `json:"ip"`
}
type vaptchaRsp struct {
	Success int    `json:"success"`
	Score   int    `json:"score"`
	Msg     string `json:"msg"`
}

func NewCaptchaDomain() *CaptchaDomain {
	return &CaptchaDomain{}
}

func (c *CaptchaDomain) Verify(server string, vid string, key string, token string, scene int, ip string) bool {
	// 发送http请求
	resp, err := tools.Post(server, vaptchaReq{
		Id:        vid,
		SecretKey: key,
		Token:     token,
		Ip:        ip,
	})
	if err != nil {
		logx.Error(err)
		return false
	}
	result := &vaptchaRsp{}
	err = json.Unmarshal(resp, result)
	if err != nil {
		logx.Error(err)
		return false
	}
	return result.Success == 1
}
