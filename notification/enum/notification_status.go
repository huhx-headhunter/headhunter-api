package enum

import (
	"database/sql/driver"
	"errors"
	"github.com/goccy/go-json"
)

type NotificationStatus struct {
	Code        int
	Description string
}

var (
	NotificationStatusInit    = NotificationStatus{Code: 0, Description: "初始化"}
	NotificationStatusSuccess = NotificationStatus{Code: 1, Description: "成功"}
	NotificationStatusFail    = NotificationStatus{Code: 2, Description: "失败"}
	NotificationStatusCancel  = NotificationStatus{Code: 3, Description: "取消"}
)

var notificationStatusByCode = map[int]NotificationStatus{
	NotificationStatusInit.Code:    NotificationStatusInit,
	NotificationStatusSuccess.Code: NotificationStatusSuccess,
	NotificationStatusFail.Code:    NotificationStatusFail,
	NotificationStatusCancel.Code:  NotificationStatusCancel,
}

func NotificationStatusFromValue(code int) NotificationStatus {
	return notificationStatusByCode[code]
}

func (r *NotificationStatus) Scan(value interface{}) error {
	code, ok := value.(int64)
	if !ok {
		return errors.New("invalid type for JobStatus")
	}
	*r = NotificationStatusFromValue(int(code))
	return nil
}

func (r NotificationStatus) Value() (driver.Value, error) {
	return r.Code, nil
}

func (r *NotificationStatus) UnmarshalJSON(data []byte) error {
	var code int
	if err := json.Unmarshal(data, &code); err != nil {
		return err
	}
	resourceType := NotificationStatusFromValue(code)
	*r = resourceType
	return nil
}
