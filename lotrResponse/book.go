package lotrResponse

type Book struct {
	Id   string `json:"_id"`
	Name string `json:"name"`
}

type ListBooks struct {
	BaseResponse

	Books []Book `json:"docs"`
}

type Chapter struct {
	Id   string `json:"_id"`
	Name string `json:"chapterName"`
}

type BookChapters struct {
	BaseResponse

	Chapters []Chapter `json:"docs"`
}
