package lotrResponse

type Character struct {
	Id      string `json:"_id"`
	Name    string `json:"name"`
	Height  string `json:"height"`
	Race    string `json:"race"`
	Gender  string `json:"gender"`
	Birth   string `json:"birth"`
	Spouse  string `json:"spouse"`
	Death   string `json:"death"`
	Realm   string `json:"realm"`
	Hair    string `json:"hair"`
	WikiUrl string `json:"wikiUrl"`
}

type ListCharacters struct {
	PaginationInfo

	Characters []Character `json:"docs"`
	Total      int         `json:"total"`
}
