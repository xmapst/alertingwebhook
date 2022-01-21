package dingtalk

const (
	MsgTypeText       = "text"
	MsgTypeLink       = "link"
	MsgTypeMarkdown   = "markdown"
	MsgTypeActionCard = "actionCard"
	MsgTypeFeedCard   = "feedCard"
	robotUrl          = "https://oapi.dingtalk.com/robot/send"
)

type textMessage struct {
	MsgType string     `json:"msgtype"`
	Text    textParams `json:"text"`
	At      atParams   `json:"at"`
}

type textParams struct {
	Content string `json:"content"`
}

type atParams struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	IsAtAll   bool     `json:"isAtAll,omitempty"`
}

type linkMessage struct {
	MsgType string     `json:"msgtype"`
	Link    linkParams `json:"link"`
}

type linkParams struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	MessageURL string `json:"messageUrl"`
	PicURL     string `json:"picUrl,omitempty"`
}

type markdownMessage struct {
	MsgType  string         `json:"msgtype"`
	Markdown markdownParams `json:"markdown"`
	At       atParams       `json:"at"`
}

type markdownParams struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type actionCardMessage struct {
	MsgType    string           `json:"msgtype"`
	ActionCard actionCardParams `json:"actionCard"`
}

type actionCardParams struct {
	Title          string `json:"title"`
	Text           string `json:"text"`
	SingleTitle    string `json:"singleTitle"`
	SingleURL      string `json:"singleURL"`
	BtnOrientation string `json:"btnOrientation,omitempty"`
	BTns           []BTns `json:"btns,omitempty"`
}

type BTns struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionUrl"`
}

type feedCardMessage struct {
	MsgType  string         `json:"msgtype"`
	FeedCard feedCardParams `json:"feed_card"`
}

type feedCardParams struct {
	Links []FeedCardLink `json:"links"`
}

type FeedCardLink struct {
	Title      string `json:"title"`
	MessageURL string `json:"messageUrl"`
	PicURL     string `json:"picUrl,omitempty"`
}
