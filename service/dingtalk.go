package service

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

// DingTalkMessage struct
type DingTalkMessage struct {
	Message string //消息
	Title   string // markdown标题
	Type    string // 消息类型
}

// DingTalkClient Client for sending message
type DingTalkClient struct {
	RobotURL string
}

func NewDingTalk() *DingTalkClient {
	return &DingTalkClient{
		RobotURL: viper.GetString("dingtalk.url"),
	}
}

// SendMessage send message to dingtalk
//https://developers.dingtalk.com/document/app/custom-robot-access
func (d *DingTalkClient) SendMessage(msg DingTalkMessage) error {

	var message string
	switch msg.Type {
	case "text":
		message = fmt.Sprintf(`{"msgtype": "text","text": {"content": "监控报警: %s"}}`, msg.Message)
	case "markdown":
		message = fmt.Sprintf(`{"msgtype": "markdown","markdown":{"title": 监控报警: "%s", "text": "%s"}}`, msg.Title, msg.Message)
	default:
		message = fmt.Sprintf(`{"msgtype": "text","text": {"content": "监控报警: %s"}}`, msg.Message)
	}

	client := &http.Client{}
	request, _ := http.NewRequest("POST", d.RobotURL, bytes.NewBuffer([]byte(message)))
	request.Header.Set("Content-type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("访问钉钉URL(%s) 出错了: %s", d.RobotURL, err)
	}
	if response.StatusCode != 200 {
		body, _ := ioutil.ReadAll(response.Body)
		return fmt.Errorf("访问钉钉URL(%s) 出错了: %s", d.RobotURL, string(body))
	}
	return nil
}
