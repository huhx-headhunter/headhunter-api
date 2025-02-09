package enum

import (
	"database/sql/driver"
	"errors"
	"github.com/goccy/go-json"
)

type NotificationType struct {
	Name        string `json:"value"`
	Description string `json:"description"`
}

var (
	NotificationTypeApp    = NotificationType{Name: "app", Description: "App应用"}
	NotificationTypeEmail  = NotificationType{Name: "email", Description: "邮件"}
	NotificationTypeSms    = NotificationType{Name: "sms", Description: "短信"}
	NotificationTypeSlack  = NotificationType{Name: "slack", Description: "Slack"}
	NotificationTypeFeishu = NotificationType{Name: "feishu", Description: "飞书"}
	NotificationTypeWechat = NotificationType{Name: "wechat", Description: "企微"}
)

var notificationTypeByName = map[string]NotificationType{
	NotificationTypeApp.Name: NotificationTypeApp,
}

func NotificationTypeFromValue(code string) NotificationType {
	return notificationTypeByName[code]
}

func (r *NotificationType) Scan(value interface{}) error {
	name, ok := value.(string)
	if !ok {
		return errors.New("invalid type for JobStatus")
	}
	*r = NotificationTypeFromValue(name)
	return nil
}

func (r NotificationType) Value() (driver.Value, error) {
	return r.Name, nil
}

func (r *NotificationType) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	resourceType := NotificationTypeFromValue(name)
	*r = resourceType
	return nil
}
