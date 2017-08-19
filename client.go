package jikeyoujia

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	Debug bool
	Token string

	Latitude  string
	Longitude string
}

const BaseURI = "http://jikeyoujia.cn:9009/flsi"

const UserAgent = "yxhd/3.1.1 (iPhone; iOS 10.3.3; Scale/2.00)" //yxhd即壹线互动，是吉客优家的开发商
const DefaultLatitude = "30.5347256"
const DefaultLongitude = "104.066828"

func New() *Client {
	return &Client{
		Latitude:  DefaultLatitude,
		Longitude: DefaultLongitude,
	}
}

func (client *Client) EnableDebug() {
	client.Debug = true
}

func (client *Client) DeviceUDID() string {
	latitudeSum := md5.Sum([]byte(client.Latitude))
	longitudeSum := md5.Sum([]byte(client.Longitude))

	return fmt.Sprintf("%x%x", latitudeSum, longitudeSum)
}

//通过基础响应判断是否请求失败的接口
type BaseResponser interface {
	CheckError() error
}

//每个接口都会有的基础响应
type BaseResponse struct {
	Msg     string
	Success string
}

//从基础响应中检查是否请求失败
func (resp BaseResponse) CheckError() error {
	if resp.Success != "true" {
		return fmt.Errorf("%s", resp.Msg)
	}
	return nil
}

//发起HTTP请求，会设置合适的User-Agent和Content-Type，当client.Token不为空时，也会默认带上
func (client *Client) request(method, path string, header http.Header, data url.Values, respInfo BaseResponser) error {
	req, err := http.NewRequest(method, BaseURI+path, strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("无法创建请求: %s", err)
	}

	//Fixme: howto generate this??
	//req.Header.Add("verify", "3AD4FE380B22D91F523A4756B636F52A")

	req.Header.Add("pdatype", "0")
	req.Header.Add("pdaname", "iOS")
	req.Header.Add("pdaos", "10.3.3")
	req.Header.Add("pdadate", time.Now().Format("2006-01-02 15:04:05"))

	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("deviceudid", client.DeviceUDID())

	req.Header.Add("apptype", "0")
	req.Header.Add("appversion", "3.1.1")

	req.Header.Add("latitude", client.Latitude)
	req.Header.Add("longitude", client.Longitude)

	req.Header.Add("cityid", "1")
	req.Header.Add("city", "高新南区")

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	//如果Token已经获取到了则默认添加
	if client.Token != "" {
		req.Header.Add("token", client.Token)
	}

	//外部接口提供的Header直接覆盖默认值
	if header != nil {
		for key, values := range header {
			req.Header.Del(key)
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("请求错误: %s", err)
	}
	defer resp.Body.Close()

	//如果响应是未知的，需要直接将结果打印出来，以便确定格式
	if respInfo == nil {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("读取错误: %s", err)
		}

		fmt.Println(path, "响应", string(b))
		return nil
	}

	err = json.NewDecoder(resp.Body).Decode(respInfo)
	if err != nil {
		return fmt.Errorf("读取错误: %s", err)
	}

	return respInfo.CheckError()
}
