//吉客优家App自动签到工具
package main

import (
	".."

	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "用法: %s 用户名 密码\n", os.Args[0])
		return
	}

	username := os.Args[1]
	password := os.Args[2]

	client := jikeyoujia.New()
	client.EnableDebug()

	fmt.Println("登陆")
	loginInfo, err := client.Login(username, password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	fmt.Println("签到前分数:", loginInfo.Score)

	fmt.Println("签到")
	err = client.UserSign(username)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	fmt.Println("获取签到后分数")
	detailInfo, err := client.UserDetail(username)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	fmt.Println("签到后分数:", detailInfo.Score)
}
