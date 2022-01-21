package dingtalk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Roboter 是由Robot实现的接口，可以发送多种类型的消息。
type Roboter interface {
	SendText(content string, atMobiles []string, isAtAll bool) error
	SendLink(title, text, messageURL, picURL string) error
	SendMarkdown(title, text string, atMobiles []string, isAtAll bool) error
	SendActionCard(title, text, singleTitle, singleURL, btnOrientation string, bTns []BTns) error
	SendFeedCard(links []FeedCardLink) error
}

// Robot 表示可以将消息发送到组的dingtalk自定义机器人
type Robot struct {
	Token  string
	SecKey string
}

// NewRobot 返回可以发送消息的Roboter
func NewRobot(token, secKey string) Roboter {
	return Robot{
		Token:  token,
		SecKey: secKey,
	}
}

// SendText 发送text类型的消息。
func (r Robot) SendText(content string, atMobiles []string, isAtAll bool) error {
	return r.robotSend(&textMessage{
		MsgType: MsgTypeText,
		Text: textParams{
			Content: content,
		},
		At: atParams{
			AtMobiles: atMobiles,
			IsAtAll:   isAtAll,
		},
	})
}

// SendLink 发送link类型消息。
func (r Robot) SendLink(title, text, messageURL, picURL string) error {
	return r.robotSend(&linkMessage{
		MsgType: MsgTypeLink,
		Link: linkParams{
			Title:      title,
			Text:       text,
			MessageURL: messageURL,
			PicURL:     picURL,
		},
	})
}

// SendMarkdown 发送markdown类型消息。
func (r Robot) SendMarkdown(title, text string, atMobiles []string, isAtAll bool) error {
	return r.robotSend(&markdownMessage{
		MsgType: MsgTypeMarkdown,
		Markdown: markdownParams{
			Title: title,
			Text:  text,
		},
		At: atParams{
			AtMobiles: atMobiles,
			IsAtAll:   isAtAll,
		},
	})
}

// SendActionCard 发送action card类型消息。
func (r Robot) SendActionCard(title, text, singleTitle, singleURL, btnOrientation string, bTns []BTns) error {
	return r.robotSend(&actionCardMessage{
		MsgType: MsgTypeActionCard,
		ActionCard: actionCardParams{
			Title:          title,
			Text:           text,
			SingleTitle:    singleTitle,
			SingleURL:      singleURL,
			BtnOrientation: btnOrientation,
			BTns:           bTns,
		},
	})
}

// SendFeedCard 发送feed card类型消息.
func (r Robot) SendFeedCard(links []FeedCardLink) error {
	return r.robotSend(&feedCardMessage{
		MsgType:  MsgTypeFeedCard,
		FeedCard: feedCardParams{Links: links},
	})
}

type dingResponse struct {
	Errcode int
	Errmsg  string
}

func (r Robot) robotSend(msg interface{}) error {
	m, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	value := url.Values{}
	value.Set("access_token", r.Token)
	if r.SecKey != "" {
		timestamp := time.Now().UnixNano() / 1e6
		sign := getRobotDingSin(timestamp, r.SecKey)
		value.Set("timestamp", fmt.Sprintf("%d", timestamp))
		value.Set("sign", sign)
	}
	request, err := http.NewRequest(http.MethodPost, robotUrl, bytes.NewReader(m))
	if err != nil {
		return fmt.Errorf("error request: %v", err.Error())
	}
	request.URL.RawQuery = value.Encode()
	request.Header.Add("Content-Type", "application/json;charset=utf-8")

	res, err := (&http.Client{}).Do(request)
	if err != nil {
		return fmt.Errorf("send dingTalk message failed, error: %v", err.Error())
	}
	defer func() { _ = res.Body.Close() }()

	result, err := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 || err != nil {
		return fmt.Errorf("send dingTalk message failed, %s", result)
	}

	var ret dingResponse
	err = json.Unmarshal(result, &ret)
	if err != nil || ret.Errcode != 0 {
		return fmt.Errorf("send dingTalk message failed, %s", result)
	}

	return nil
}

func getRobotDingSin(timestamp int64, secret string) (sign string) {
	strToHash := fmt.Sprintf("%d\n%s", timestamp, secret)
	hmac256 := hmac.New(sha256.New, []byte(secret))
	hmac256.Write([]byte(strToHash))
	data := hmac256.Sum(nil)
	return base64.StdEncoding.EncodeToString(data)
}
