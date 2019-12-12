//吉客优家App自动签到工具
package main

import (
	"jikeyoujia"

	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

// 腾讯云函数触发事件
type ScfEvent struct {
	Message string
}

func ScfHandle(ctx context.Context, e ScfEvent) error {
	var param map[string]string
	err := json.Unmarshal([]byte(e.Message), &param)
	if err != nil {
		return fmt.Errorf("传入参数不合法: %s", err)
	}

	return doSign(param["user"], param["pass"])
}

func LambdaHandle() error {
	return doSign(os.Getenv("USER"), os.Getenv("PASS"))
}

func doSign(username, password string) error {
	client := jikeyoujia.New()
	client.EnableDebug()

	fmt.Println("登陆用户", username)
	loginInfo, err := client.Login(username, password)
	if err != nil {
		return err
	}

	fmt.Printf("登陆信息: %+v", loginInfo)
	fmt.Println("签到前分数:", loginInfo.Score)

	fmt.Println("获取签到清单")
	scoreList, err := client.UserScoreList(username)
	if err != nil {
		return err
	}
	fmt.Printf("签到清单信息: %+v", scoreList)

	fmt.Println("签到")
	signInfo, err := client.UserSign(username)
	if err != nil {
		return err
	}

	fmt.Printf("签到信息: %+v", signInfo)

	fmt.Println("获取签到后分数")
	detailInfo, err := client.UserDetail(username)
	if err != nil {
		return err
	}

	fmt.Printf("签到详情信息: %+v", detailInfo)
	fmt.Println("签到后分数:", detailInfo.Score)

	return nil
}

func main() {
	//检测是否运行于腾讯云函数环境
	_, ok := os.LookupEnv("TENCENTCLOUD_RUNENV")
	if ok {
		cloudfunction.Start(ScfHandle)
		return
	}

	_, ok = os.LookupEnv("AWS_LAMBDA_RUNTIME_API")
	if ok {
		lambda.Start(LambdaHandle)
	}

	//其他情况直接从命令行获取参数运行
	if len(os.Args) < 3 {
		fmt.Printf("用法: %s 用户名 密码\n", os.Args[0])
		return
	}

	err := doSign(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
}
