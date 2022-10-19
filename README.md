# The Lord of the Rings SDK

This is an SDK for the [LOTR API](https://the-one-api.dev/documentation)

## Installation

You can use "go get" to install the SDK

```bash
go get -u github.com/nishant4/nishant_varshney-SDK
```

## Usage
All the main available functions are below. Please check the API documentations to better understand.

```go
import "lotr"

app := lotr.New("Access Token")

// No Token required
//Books
ListBooks() // List of all "The Lord of the Rings" books
GetBook(bookId) // Request one specific book
GetBookChapters(bookId) // Request all chapters of one specific book

// Token Required
// Movies
ListMovies() // List of all movies
GetMovie(movieId) // Request one specific movie
GetMovieQuotes(movieId) // Request all movie quotes for one specific movie

// Characters
ListCharacters() // List of characters including metadata
GetCharacter(characterId) // Request one specific character
GetCharacterQuotes(characterId) // Request all movie quotes of one specific character

// Quotes
ListQuotes() // List of all movie quotes
GetQuote(quoteId) // Request one specific movie quote

// Chapters
ListChapters() // List of all book chapters
GetChapter(chapterId) // Request one specific book chapter
```

### Response
Responses are defined in lotr/lotrResponse

```go
type BaseResponse struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Page   int `json:"page"`
	Pages  int `json:"pages"`

	Total int `json:"total"`
}

// Books Response
type Book struct {
	Id   string `json:"_id"`
	Name string `json:"name"`
}

type ListBooks struct {
	BaseResponse // inherit

	Books []Book `json:"docs"`
}

// Others
ListCharacters
ListMovies
ListQuotes
ListChapters

// Detailed structures
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

type Quote struct {
	Id        string `json:"_id"`
	Dialog    string `json:"dialog"`
	Movie     string `json:"movie"`
	Character string `json:"Character"`
}

type Chapter struct {
	Id   string `json:"_id"`
	Name string `json:"chapterName"`
	Book string `json:"book"`
}
```

### Control Output

You can use GetOptions to control the output. 

Every function that returns more than one result such as GetMovieQuotes has a counterpart function that takes GetOptions as an extra parameter.

Ex - GetMovieQuotesOptions(movieId, options)

```go
options := lotr.NewGetOptions()

// other shortcuts
NewGetOptionsOffset(offset)
NewGetOptionsPage(pageNumber)
NewGetOptionsPageLimited(pageNumber, limit)

type GetOptions struct {
	Limit  int
	Offset int
	Page   int

	Match       map[string][]string // param -> values. Ex - race -> Elf,Human
	NotMatch    map[string][]string
	GreaterThan map[string]int // key > value
	LessThan    map[string]int // key < value

	SortKey string // key to sort on
	SortAsc bool // order for the sort, defaults false
}
```

### Pagination
Use Limit, Offset, and Page in GetOptions for Pagination. Offset will override page if > 0


### Filtering
#### Match
You can use any of the output parameters as key but the API doesn't support all of them. Use the array to give a list of possible values.

You can use regex here as well. Just use the format "/regex/i" as the value

Ex - name = Gandalf,Galadriel

#### Not Match
Ex - name != Gandalf, Galadriel

#### GreaterThan & Less Than
Ex - runtimeInMinutes > 150, runtimeInMinutes < 250

#### Sorting
Use options.SortOnKey(key, asc)

## Testing

All major functions have basic tests in the test folder. Usage -

```go
go test ./test -v
```

## Other Important Notes
1. All the endpoints are defined in config.go file. They can be overridden in case the external API changes
2. For the calls that don't require an access token, just create App with an empty token
3. You can get the raw output of any endpoint by calling app.ApiOptions(options)
4. Match/NotMatch can be used to mimic Include/Exclude functionality
5. Exists/NotExists functionality was not working for the API 
6. Currently the flow of code is logged to stdout. It can be disabled by calling log.SetOutput(ioutil.Discard)


## Possible Improvements
1. Didn't feel the need for a design.md file. Readme explains most of the code in detail. The code should be self esplantory in itself. It is verbose enough and divided into proper modules.
2. Routes should be moved to their own folder
3. Tests should be implemented without hardcoding values as the API might change
