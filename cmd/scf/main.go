//吉客优家App自动签到工具
package main

import (
	"context"
	"log"
	"os"

	"jikeyoujia"

	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

func main() {
	cloudfunction.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, event interface{}) (string, error) {
	username := os.Getenv("JIKE_USER")
	password := os.Getenv("JIKE_PASS")

	client := jikeyoujia.New()
	client.EnableDebug()

	log.Println("登陆用户", username)
	loginInfo, err := client.Login(username, password)
	if err != nil {
		return "", err
	}

	log.Printf("登陆信息: %+v", loginInfo)
	log.Println("签到前分数:", loginInfo.Score)

	log.Println("获取签到清单")
	scoreList, err := client.UserScoreList(username)
	if err != nil {
		return "", err
	}
	log.Printf("签到清单信息: %+v", scoreList)

	log.Println("签到")
	signInfo, err := client.UserSign(username)
	if err != nil {
		return "", err
	}

	log.Printf("签到信息: %+v", signInfo)

	log.Println("获取签到后分数")
	detailInfo, err := client.UserDetail(username)
	if err != nil {
		return "", err
	}

	log.Printf("签到详情信息: %+v", detailInfo)
	log.Println("签到后分数:", detailInfo.Score)

	return "", nil
}
