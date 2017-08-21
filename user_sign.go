package jikeyoujia

import (
	"encoding/json"
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
func (client *Client) UserSign(username string) error {
	hdr := http.Header{}
	hdr.Set("username", username)

	var info SignResponse

	err := client.request("POST", "/users!usersign.action", hdr, nil, &info)
	if err != nil {
		return fmt.Errorf("请求错误: %s", err)
	}

	if client.Debug {
		b, _ := json.MarshalIndent(info, "", "\t")
		fmt.Println("签到响应：\n", string(b))
	}

	return nil
}
