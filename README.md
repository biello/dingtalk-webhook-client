# dingtalk-webhook-client

## usage

```
import "github.com/biello/dingtalk-webhook-client/client"


func main() {

	// change param to your webhook
	cli := client.DefaultDingTalkClient("https://oapi.dingtalk.com/robot/send?access_token=xxxxx")

	// text request
	req := client.CreateOapiRobotSendTextRequest("我就是我，是不一样的烟火", []string{"158xxxxxxxx", "137xxxxxxxx"}, false)

	// execute
	_, err := cli.Execute(req)
	if err != nil {
		fmt.Printf("send fail:%s", err)
	}

}

```
