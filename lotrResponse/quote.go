package lotrResponse

type Quote struct {
	Id        string `json:"_id"`
	Dialog    string `json:"dialog"`
	Movie     string `json:"movie"`
	Character string `json:"Character"`
}

type ListQuotes struct {
	BaseResponse

	Quotes []Quote `json:"docs"`
}
