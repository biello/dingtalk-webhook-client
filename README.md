# dingtalk-webhook-client

## usage

```
  
req := CreateOapiRobotSendTextRequest("我就是我，是不一样的烟火", nil, false)
resp, err := DefaultDingTalkClient(conf.Conf.DingTalk.Webhooks[0]).Execute(req)
if err != nil || resp.ErrCode != 0 {
	t.Logf("send dingtalk msg err: %s，resp: %v", err, resp)
}


```
