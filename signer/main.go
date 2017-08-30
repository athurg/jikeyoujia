//吉客优家App自动签到工具
package main

import (
	".."

	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatalf("用法: %s 用户名 密码 微信通知用户\n", os.Args[0])
		return
	}

	username := os.Args[1]
	password := os.Args[2]
	toUser := os.Args[3]

	client := jikeyoujia.New()
	client.EnableDebug()

	log.Println("登陆")
	loginInfo, err := client.Login(username, password)
	if err != nil {
		Notify(toUser, "吉客优家登陆失败: "+err.Error())
		log.Fatalf("%s\n", err)
	}

	log.Printf("登陆信息: %+v", loginInfo)
	log.Println("签到前分数:", loginInfo.Score)

	log.Println("获取签到清单")
	scoreList, err := client.UserScoreList(username)
	if err != nil {
		Notify(toUser, "吉客优家获取签到单失败: "+err.Error())
		log.Fatalf("%s\n", err)
	}
	log.Printf("签到清单信息: %+v", scoreList)

	log.Println("签到")
	signInfo, err := client.UserSign(username)
	if err != nil {
		Notify(toUser, "吉客优家签到失败: "+err.Error())
		log.Fatalf("%s\n", err)
	}

	log.Printf("签到信息: %+v", signInfo)

	log.Println("获取签到后分数")
	detailInfo, err := client.UserDetail(username)
	if err != nil {
		Notify(toUser, "吉客优家获取签到后分数失败: "+err.Error())
		log.Fatalf("%s\n", err)
	}

	log.Printf("签到详情信息: %+v", detailInfo)
	log.Println("签到后分数:", detailInfo.Score)

	Notify(toUser, "吉客优家签到成功, 分数变化"+loginInfo.Score+" => "+detailInfo.Score)
}

func Notify(toUser, message string) {
	log.Println("发送通知", message)
	getParams := url.Values{}
	getParams.Set("touser", toUser)
	getParams.Set("content", message)

	_, err := http.Get("http://localhost:8083/default_notify?" + getParams.Encode())
	if err != nil {
		log.Println("微信通知失败", err)
	}
}
