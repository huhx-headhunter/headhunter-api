package resp

import (
	"github.com/huhx-headhunter/headhunter-api/notification/enum"
	"github.com/huhx/common-go/times"
)

type NotificationResponse struct {
	Id         string                   `json:"id" example:"1800783867820785152"`
	Title      string                   `json:"title"`
	Content    string                   `json:"content"`
	Receiver   string                   `json:"receiver"`
	Source     enum.NotificationSource  `json:"source"`
	Type       string                   `json:"type"`
	ActiveTime times.LocalDateTime      `json:"activeTime" example:"2019-10-03T12:00:00"`
	Metadata   map[string]interface{}   `json:"metadata"`
	Records    []NotificationRecordItem `json:"records"`
}

type NotificationRecordItem struct {
	Id             int64               `json:"id"`
	NotificationId int64               `json:"notificationId"`
	Status         int                 `json:"status"`
	Reason         string              `json:"reason"`
	CreatedAt      times.LocalDateTime `json:"createdAt" example:"2019-10-03T12:00:00"`
}
