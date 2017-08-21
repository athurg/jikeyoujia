package jikeyoujia

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/url"
)

type UserLoginResponse struct {
	BaseResponse

	Token string

	Accamt string

	Ordername    string
	Ordercity    string
	Orderarea    string
	Orderaddress string
	Ordermobile  string

	Homeactivity string
	Score        string
	Timerange    string
	Platenumber  string

	City   string
	Cityid string

	Name          string
	Username      string
	Email         string
	Userstatus    string
	Gender        string
	Address       string
	Contactmobile string

	Grade          string
	Graderemark    string
	Graderemarkurl string

	Carserviceid string

	Lat string
	Lng string

	Ismillionfamily string
	Headimg         string

	Worktimebegin string
	Worktimeend   string

	Couponcount string

	Dispatchmid string
	Dispatchall string

	UnreadAnswerreply string `json:"unread_answerreply"`
}

func (client *Client) Login(username, password string) (*UserLoginResponse, error) {
	var info UserLoginResponse

	data := url.Values{}
	data.Set("username", username)
	data.Set("pwd", fmt.Sprintf("%X", md5.Sum([]byte(password))))

	err := client.request("POST", "/users!loginverify.action", nil, data, &info)
	if err != nil {
		return nil, fmt.Errorf("请求错误:%s", err)
	}

	if client.Debug {
		b, _ := json.MarshalIndent(info, "", "\t")
		fmt.Println("登陆响应：\n", string(b))
	}

	client.Token = info.Token

	return &info, nil
}
