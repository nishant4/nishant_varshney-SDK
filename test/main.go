package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"lotr"
)

const ACCESS_CODE = "mRNnJRcQoHPw3duyL6Vp"

func prettyPrint(jsonBody []byte) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, jsonBody, "", "\t")
	if err != nil {
		log.Println("PrettyPrint: error: ", err)
		return
	}

	log.Println(string(prettyJSON.Bytes()))
}

func prettyPrintStruct(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	log.Println(string(s))
}

func main() {
	app := lotr.New("mRNnJRcQoHPw3duyL6Vp")
	fmt.Println("Created new")

	// s := app.ListBooks()
	// prettyPrint(s)

	// prettyPrint(app.GetBook("5cf5805fb53e011a64671582"))
	// prettyPrint(app.GetBookChapters("5cf5805fb53e011a64671582"))

	// bookList, _ := app.ListBooks()
	// prettyPrintStruct(bookList)

	// book, _ := app.GetBook(bookList.Books[0].Id)
	// prettyPrintStruct(book)

	// bookChapters, _ := app.GetBookChapters(bookList.Books[0].Id)
	// prettyPrintStruct(bookChapters)

	// characters, _ := app.ListCharacters()
	// prettyPrintStruct(characters)

	options := lotr.NewGetOptions()
	options.Match = map[string][]string{
		"name": []string{"/Adrahil/i", "Gandalf", "Aegnor", "Aerin"},
	}
	options.Limit = 2
	options.Page = 3
	characters, _ := app.ListCharactersOptions(options)
	prettyPrintStruct(characters)

	// characters, _ = app.ListCharactersOptions(lotr.NewGetOptionsOffsetLimited(0, 5))
	// prettyPrintStruct(characters)

	// characters, _ := app.ListCharactersOffsetWithOptions(options, 1, -1)
	// prettyPrintStruct(characters)

	// characters, _ := app.ListCharactersPaginatedWithOptions(options, 2, 2)
	// prettyPrintStruct(characters)
}
