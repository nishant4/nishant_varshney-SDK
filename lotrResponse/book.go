package lotrResponse

type Book struct {
	Id   string `json:"_id"`
	Name string `json:"name"`
}

type ListBooks struct {
	PaginationInfo

	Books []Book `json:"docs"`
	Total int    `json:"total"`
}

type GetBook struct {
	PaginationInfo

	Book []Book `json:"docs"`
}

type Chapter struct {
	Id   string `json:"_id"`
	Name string `json:"chapterName"`
}

type GetBookChapters struct {
	PaginationInfo

	Chapters []Chapter `json:"docs"`
	Total    int       `json:"total"`
}
