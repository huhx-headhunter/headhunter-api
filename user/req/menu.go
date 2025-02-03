package req

type MenuAddRequest struct {
	Channel   string `json:"channel"`
	Parent    int64  `json:"parent"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	Component string `json:"component"`
}

type MenuUpdateRequest struct {
	Id        int64  `json:"id"`
	Channel   string `json:"channel"`
	Parent    int64  `json:"parent"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	Component string `json:"component"`
}
