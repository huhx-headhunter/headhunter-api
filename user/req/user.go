package req

import "github.com/huhx/common-go/times"

type UserAddRequest struct {
	Name      string          `json:"name"`
	Username  string          `json:"username"`
	Password  string          `json:"password"`
	Nickname  string          `json:"nickname"`
	Signature string          `json:"signature"`
	Avatar    string          `json:"avatar"`
	Birthday  times.LocalDate `json:"birthday"`
	Address   string          `json:"address"`
	Email     string          `json:"email"`
	Phone     string          `json:"phone"`
	Wechat    string          `json:"wechat"`
	RoleId    int64           `json:"roleId"` // boss, emp, admin
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateRequest struct {
	Id          int64           `json:"id"`
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
