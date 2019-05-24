package main

import (
	"fmt"
	"github.com/biello/dingtalk-webhook-client/client"
)

func main() {

	// change param to your webhook
	cli := client.DefaultDingTalkClient("https://oapi.dingtalk.com/robot/send?access_token=xxxxx")

	// text request
	req := client.CreateOapiRobotSendTextRequest("some msg...", []string{"158xxxxxxxx", "137xxxxxxxx"}, false)

	// execute
	_, err := cli.Execute(req)
	if err != nil {
		fmt.Printf("send fail:%s", err)
	}

}
