package resp

import "github.com/huhx/common-go/times"

type TokenResponse struct {
	Token string `json:"token"`
}

type UserListRespItem struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type UserDetailResp struct {
	Id          string          `json:"id"`
	Name        string          `json:"name"`
	Username    string          `json:"username"`
	Password    string          `json:"password"`
	Nickname    string          `json:"nickname"`
	Avatar      string          `json:"avatar"`
	Birthday    times.LocalDate `json:"birthday"`
	Address     string          `json:"address"`
	Signature   string          `json:"signature"`
	Email       string          `json:"email"`
	Phone       string          `json:"phone"`
	Wechat      string          `json:"wechat"`
	RoleId      int64           `json:"roleId"` // boss, emp, admin
	FailedCount int             `json:"failedCount"`
}
