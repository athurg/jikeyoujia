//吉客优家App自动签到工具
package main

import (
	".."

	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("用法: %s 用户名 密码\n", os.Args[0])
		return
	}

	username := os.Args[1]
	password := os.Args[2]

	client := jikeyoujia.New()
	client.EnableDebug()

	log.Println("登陆")
	loginInfo, err := client.Login(username, password)
	if err != nil {
		log.Fatalf("%s\n", err)
		return
	}
	log.Printf("登陆信息: %+v", loginInfo)
	log.Println("签到前分数:", loginInfo.Score)

	log.Println("获取签到清单")
	scoreList, err := client.UserScoreList(username)
	if err != nil {
		log.Fatalf("%s\n", err)
		return
	}
	log.Printf("签到清单信息: %+v", scoreList)

	log.Println("签到")
	signInfo, err := client.UserSign(username)
	if err != nil {
		log.Fatalf("%s\n", err)
		return
	}
	log.Printf("签到信息: %+v", signInfo)

	log.Println("获取签到后分数")
	detailInfo, err := client.UserDetail(username)
	if err != nil {
		log.Fatalf("%s\n", err)
		return
	}

	log.Printf("签到详情信息: %+v", detailInfo)
	log.Println("签到后分数:", detailInfo.Score)
}
