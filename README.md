# dingtalk-webhook-client
Go语言钉钉群机器人客户端，支持发送文本、链接卡片、Markdown类型的消息
### 文本
![文本](https://img.alicdn.com/tfs/TB1jFpqaRxRMKJjy0FdXXaifFXa-497-133.png#align=left&display=inline&height=112&originHeight=133&originWidth=497&status=done&width=418)
### 链接卡片
![链接卡片](https://img.alicdn.com/tfs/TB1VfZtaUgQMeJjy0FeXXXOEVXa-498-193.png#align=left&display=inline&height=138&originHeight=193&originWidth=498&status=done&width=355)
### Markdown
![Markdown](https://img.alicdn.com/tfs/TB1yL3taUgQMeJjy0FeXXXOEVXa-492-380.png#align=left&display=inline&height=241&originHeight=380&originWidth=492&status=done&width=312)

钉钉群机器人官方介绍：https://open.dingtalk.com/document/orgapp/robot-overview

## usage

```
import (
	"fmt"
	"github.com/biello/dingtalk-webhook-client/client"
)

// 钉钉群机器人 webhook 和 secret, 未开启加签 secret 留空
const (
	webhook = "https://oapi.dingtalk.com/robot/send?access_token=xxx"
	secret  = ""
)

func main() {

	cli := client.DefaultDingTalkClient(webhook, secret)

	// text message
	textReq := client.CreateOapiRobotSendTextRequest("我就是我, 是不一样的烟火@156xxxx8827", []string{"156xxxx8827", "189xxxx8325"}, false)

	_, err := cli.Execute(textReq)
	if err != nil {
		fmt.Printf("send fail:%s\n", err)
	}

	// link message
	linkReq := client.OapiRobotSendRequest{
		MsgType: "link",
		Link: client.Link{
			Title:      "时代的火车向前开",
			Text:       "这个即将发布的新版本，创始人陈航（花名“无招”）称它为“红树林”。而在此之前，每当面临重大升级，产品经理们都会取一个应景的代号，这一次，为什么是“红树林”？",
			PicUrl:     "",
			MessageUrl: "https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI",
		},
	}
	_, err = cli.Execute(linkReq)
	if err != nil {
		fmt.Printf("send fail:%s\n", err)
	}

	// markdown message
	markDownReq := client.OapiRobotSendRequest{
		MsgType: "markdown",
		Markdown: client.Markdown{
			Title: "杭州天气",
			Text: "#### 杭州天气 @156xxxx8827\n" +
				"> 9度，西北风1级，空气良89，相对温度73%\n\n" +
				"> ![screenshot](https://gw.alipayobjects.com/zos/skylark-tools/public/files/84111bbeba74743d2771ed4f062d1f25.png)\n" +
				"> ###### 10点20分发布 [天气](http://www.thinkpage.cn/) \n",
		},
		At: client.At{
			AtMobiles: []string{"156xxxx8827", "189xxxx8325"},
			IsAtAll:   false,
		},
	}
	_, err = cli.Execute(markDownReq)
	if err != nil {
		fmt.Printf("send fail:%s\n", err)
	}
}


```
