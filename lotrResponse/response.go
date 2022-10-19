package lotrResponse

type BaseResponse struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Page   int `json:"page"`
	Pages  int `json:"pages"`

	Total int `json:"total"`
}
