package jikeyoujia

import (
	"fmt"
	"net/http"
	"net/url"
)

type UserScoreListResponse struct {
	BaseResponse

	PageCount int
	PageIndex int
	Total     int

	Data []struct {
		Date   string
		Remark string
		Score  int
	}
}

func (client *Client) UserScoreList(username string) (*UserScoreListResponse, error) {
	var info UserScoreListResponse

	hdr := http.Header{}
	hdr.Set("username", username)

	data := url.Values{}
	data.Set("pagesize", "100")
	data.Set("pageindex", "1")

	err := client.request("POST", "/users!scorelist.action", hdr, data, &info)
	if err != nil {
		return nil, fmt.Errorf("请求错误:%s", err)
	}

	return &info, nil
}
