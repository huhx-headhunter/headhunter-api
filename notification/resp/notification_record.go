package resp

import "github.com/huhx/common-go/times"

type NotificationRecordResponse struct {
	Id             int64               `json:"id"`
	NotificationId int64               `json:"notificationId"`
	Status         int                 `json:"status"`
	Reason         string              `json:"reason"`
	CreatedAt      times.LocalDateTime `json:"createdAt" example:"2019-10-03T12:00:00"`
}
