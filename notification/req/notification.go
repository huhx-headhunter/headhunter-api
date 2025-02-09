package req

import (
	"github.com/huhx-headhunter/headhunter-api/notification/enum"
	"github.com/huhx/common-go/times"
)

type NotificationAddRequest struct {
	Title      string                  `json:"title"`
	Content    string                  `json:"content"`
	Receiver   string                  `json:"receiver"`
	Source     enum.NotificationSource `json:"source"`
	Type       string                  `json:"type"`
	ActiveTime times.LocalDateTime     `json:"activeTime" example:"2019-10-03T12:00:00"`
	Metadata   *map[string]interface{} `json:"metadata"`
}
