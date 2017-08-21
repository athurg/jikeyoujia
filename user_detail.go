package jikeyoujia

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserDetailResponse struct {
	BaseResponse

	Accamt             string
	Contactmobile      string
	Homeactivity       string
	Score              string
	Timerange          string
	Platenumber        string
	Username           string
	Token              string
	Name               string
	Gender             string
	Grade              string
	Carserviceid       string
	Ismillionfamily    string
	Headimg            string
	Worktimebegin      string
	Couponcount        string
	Graderemarkurl     string
	Worktimeend        string
	Graderemark        string
	Dispatchmid        string
	Address            string
	Userstatus         string
	Email              string
	Dispatchall        string
	Unread_answerreply string
}

func (client *Client) UserDetail(username string) (*UserDetailResponse, error) {
	var info UserDetailResponse

	hdr := http.Header{}
	hdr.Set("username", username)
	err := client.request("POST", "/users!detail.action", hdr, nil, &info)
	if err != nil {
		return nil, fmt.Errorf("请求错误:%s", err)
	}

	if client.Debug {
		b, _ := json.MarshalIndent(info, "", "\t")
		fmt.Println("用户详情响应：\n", string(b))
	}

	return &info, nil
}
