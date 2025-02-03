package req

type RoleAddRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RoleUpdateRequest struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
