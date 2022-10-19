package lotrResponse

type Chapter struct {
	Id   string `json:"_id"`
	Name string `json:"chapterName"`
	Book string `json:"book"`
}

type ListChapters struct {
	BaseResponse

	Chapters []Chapter `json:"docs"`
}
