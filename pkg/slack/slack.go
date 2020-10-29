package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/vdgonc/send2channel/pkg/config"
)

// Fields - fields of message
type Fields struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

// Attachment - attachment
type Attachment struct {
	Fallback   string   `json:"fallback"`
	Color      string   `json:"color"`
	AuthorName string   `json:"author_name"`
	Title      string   `json:"title"`
	TitleLink  string   `json:"title_link"`
	Text       string   `json:"text"`
	Footer     string   `json:"footer"`
	Fields     []Fields `json:"fields"`
}

type Message struct {
	Attachments []Attachment `json:"attachments"`
}

func timerformat() string {
	t := time.Now()

	return fmt.Sprintf("%d/%d/%d %d:%d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())
}

// newMessage - return message type
func newMessage(appname, status string) Message {
	cf := config.New()

	var fallback string
	var color string
	var text string

	switch status {
	case "init":
		fallback = cf.Fallback.Init
		color = cf.Color.Init
		text = cf.Message.Init
	case "success":
		fallback = cf.Fallback.Success
		color = cf.Color.Success
		text = cf.Message.Success
	case "fail":
		fallback = cf.Fallback.Fail
		color = cf.Color.Fail
		text = cf.Message.Fail
	default:
		fallback = appname
		color = "#ddd9d9"
		text = ""

	}

	msg := Message{
		Attachments: []Attachment{
			{
				Fallback:   fallback,
				Color:      color,
				AuthorName: appname,
				Title:      cf.Title + appname,
				Text:       text,
				Fields: []Fields{{
					Title: "Date",
					Value: timerformat(),
					Short: false,
				},
					{
						Title: "Status",
						Value: status,
						Short: false,
					},
				},
				Footer: cf.Footer,
			},
		},
	}

	return msg
}

// SendMessage - send message to channel
func SendMessage(url, appname, status string, ch chan string) {
	msg := newMessage(appname, status)

	j, _ := json.Marshal(&msg)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(j))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	txt := fmt.Sprintf("[+] Notification Sent")
	ch <- txt
}
