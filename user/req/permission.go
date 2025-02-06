package req

type PermissionAddRequest struct {
	Name        string `json:"name"`
	Code        string `json:"code"` // user:create
	Description string `json:"description"`
}
