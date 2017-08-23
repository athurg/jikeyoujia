package jikeyoujia

import (
	"fmt"
	"net/http"
)

type SignResponse struct {
	BaseResponse

	GetScore     string `json:"getscore"` //本次签到获取的积分
	SignScore    string `json:"getscore"` //签到累计获得的积分
	CurrentScore string `json:"getscore"` //当前积分
}

//签到
func (client *Client) UserSign(username string) (*SignResponse, error) {
	hdr := http.Header{}
	hdr.Set("username", username)

	var info SignResponse

	err := client.request("POST", "/users!usersign.action", hdr, nil, &info)
	if err != nil {
		return nil, fmt.Errorf("请求错误: %s", err)
	}

	return &info, nil
}
