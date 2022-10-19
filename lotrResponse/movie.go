package lotrResponse

type Movie struct {
	Id                         string  `json:"_id"`
	Name                       string  `json:"name"`
	RuntimeInMinutes           int     `json:"runtimeInMinutes"`
	BudgetInMillions           float64 `json:"budgetInMillions"`
	BoxOfficeRevenueInMillions float64 `json:"boxOfficeRevenueInMillions"`
	AcademyAwardNominations    int     `json:"academyAwardNominations"`
	AcademyAwardWins           int     `json:"academyAwardWins"`
	RottenTomatoesScore        float64 `json:"rottenTomatoesScore"`
}

type ListMovies struct {
	BaseResponse

	Movies []Movie `json:"docs"`
}
