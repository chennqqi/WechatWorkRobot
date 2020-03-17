package WechatWorkRobot

import (
	"encoding/json"
	"fmt"

	"github.com/unliar/utils/go/http"
)

// 详情 refer https://work.weixin.qq.com/api/doc/90000/90136/91770

type Robot struct {
	Key string
}

// RobotResponse 机器人接口响应
type RobotResponse struct {
	ErrorCode    int64  `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}

// NewsItem 图文消息item
type NewsItem struct {
	Title       string `json:"title"`                 // 标题，不超过128个字节，超过会自动截断
	Description string `json:"description,omitempty"` // 描述，不超过512个字节，超过会自动截断
	URL         string `json:"url"`                   // 点击后跳转的链接。
	Picurl      string `json:"picurl,omitempty"`      // 图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图 1068*455，小图150*150。
}

// RobotRequest 消息请求体
type RobotRequest struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text,omitempty"`
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown,omitempty"`
	Image struct {
		Base64 string `json:"base64"`
		MD5    string `json:"md5"`
	} `json:"image,omitempty"`
	News struct {
		Articles []*NewsItem `json:"articles"`
	} `json:"news,omitempty"`
}

// CreateBaseURL 拼接地址
func (r *Robot) CreateBaseURL() string {
	return fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", r.Key)
}

// SendText 发送纯文本
func (r *Robot) SendText(text string) (res *RobotResponse, err error) {
	data := RobotRequest{
		MsgType: "text",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: text,
		},
	}
	baseURL := r.CreateBaseURL()
	b, err := http.Post(baseURL, data, nil)
	if err != nil {
		return &RobotResponse{
			ErrorCode:    -1,
			ErrorMessage: "请求出错",
		}, err
	}
	_ = json.Unmarshal(b, &res)
	return
}

// SendMarkdown 发送markdown
func (r *Robot) SendMarkdown(markdown string) (res *RobotResponse, err error) {
	data := RobotRequest{
		MsgType: "markdown",
		Markdown: struct {
			Content string `json:"content"`
		}{
			Content: markdown,
		},
	}
	baseURL := r.CreateBaseURL()
	b, err := http.Post(baseURL, data, nil)
	if err != nil {
		return &RobotResponse{
			ErrorCode:    -1,
			ErrorMessage: "请求出错",
		}, err
	}
	_ = json.Unmarshal(b, &res)
	return
}

// SendImage 发送图片
func (r *Robot) SendImage(base64, md5 string) (res *RobotResponse, err error) {
	data := RobotRequest{
		MsgType: "image",
		Image: struct {
			Base64 string `json:"base64"`
			MD5    string `json:"md5"`
		}{
			Base64: base64,
			MD5:    md5,
		},
	}
	baseURL := r.CreateBaseURL()
	b, err := http.Post(baseURL, data, nil)
	if err != nil {
		return &RobotResponse{
			ErrorCode:    -1,
			ErrorMessage: "请求出错",
		}, err
	}
	_ = json.Unmarshal(b, &res)
	return
}

// SendNews 发送图文
func (r *Robot) SendNews(news []*NewsItem) (res *RobotResponse, err error) {
	data := RobotRequest{
		MsgType: "news",
		News: struct {
			Articles []*NewsItem `json:"articles"`
		}{
			Articles: news,
		},
	}
	baseURL := r.CreateBaseURL()
	b, err := http.Post(baseURL, data, nil)
	if err != nil {
		return &RobotResponse{
			ErrorCode:    -1,
			ErrorMessage: "请求出错",
		}, err
	}
	_ = json.Unmarshal(b, &res)
	return
}
