package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type DingTalkClient struct {
	webhook string
	secret  string
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

func DefaultDingTalkClient(webhook, secret string) DingTalkClient {
	return DingTalkClient{
		webhook: webhook,
		secret:  secret,
	}
}

func CreateOapiRobotSendTextRequest(content string, atMobiles []string, isAtAll bool) OapiRobotSendRequest {
	return OapiRobotSendRequest{
		MsgType: "text",
		Text:    Text{Content: content},
		At:      At{AtMobiles: atMobiles, IsAtAll: isAtAll},
	}

}

func (d DingTalkClient) Execute(request OapiRobotSendRequest) (*OapiRobotSendResponse, error) {
	reqBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	pushUrl, err := d.getPushUrl()
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(pushUrl, "application/json", bytes.NewReader(reqBytes))
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

func (d DingTalkClient) getPushUrl() (string, error) {
	if d.secret == "" {
		return d.webhook, nil
	}

	timestamp := strconv.FormatInt(time.Now().Unix()*1000, 10)
	sign, err := d.sign(timestamp)
	if err != nil {
		return d.webhook, err
	}

	query := url.Values{}
	query.Set("timestamp", timestamp)
	query.Set("sign", sign)
	return d.webhook + "&" + query.Encode(), nil
}

func (d DingTalkClient) sign(timestamp string) (string, error) {
	stringToSign := fmt.Sprintf("%s\n%s", timestamp, d.secret)
	h := hmac.New(sha256.New, []byte(d.secret))
	if _, err := io.WriteString(h, stringToSign); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}
