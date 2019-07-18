package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DingTalkClient struct {
	webhook string
}

type OapiRobotSendRequest struct {
	MsgType  string   `json:"msgtype"`
	Text     Text     `json:"text"`
	Markdown Markdown `json:"markdown"`
	Link     Link     `json:"link"`
	At       At       `json:"at"`
}

type Text struct {
	Content string `json:"content"`
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type Link struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	PicUrl     string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

type OapiRobotSendResponse struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode int64  `json:"errcode"`
}

func DefaultDingTalkClient(webhook string) DingTalkClient {
	return DingTalkClient{
		webhook: webhook,
	}
}

func CreateOapiRobotSendTextRequest(content string, atMobiles []string, isAtAll bool) OapiRobotSendRequest {
	return OapiRobotSendRequest{
		MsgType: "text",
		Text:    Text{Content: content},
		At:      At{AtMobiles: atMobiles, IsAtAll: isAtAll},
	}

}

func (c DingTalkClient) Execute(request OapiRobotSendRequest) (*OapiRobotSendResponse, error) {
	reqBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.webhook, "application/json", bytes.NewReader(reqBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("http response status code is: %d", resp.StatusCode))
	}
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var oResponse OapiRobotSendResponse
	if err = json.Unmarshal(responseBody, &oResponse); err != nil {
		return nil, err
	}
	if oResponse.ErrCode != 0 {
		return &oResponse, errors.New(fmt.Sprintf("response: %s", responseBody))
	}
	return &oResponse, nil
}
