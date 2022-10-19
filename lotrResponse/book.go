package lotrResponse

type Book struct {
	Id   string `json:"_id"`
	Name string `json:"name"`
}

type ListBooks struct {
	BaseResponse

	Books []Book `json:"docs"`
}
