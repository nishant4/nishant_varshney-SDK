package main

import (
	"io/ioutil"
	"log"
	"lotr"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	code := m.Run()
	os.Exit(code)
}

func TestBooks(t *testing.T) {
	app := lotr.New("")

	books, err := app.ListBooks()
	if err != nil {
		t.Error("Error ListBooks ", err)
	}

	if books.Total != 3 {
		t.Error("Should have 3 books ", books)
	}

	bookId := books.Books[0].Id
	book, err := app.GetBook(bookId)
	if err != nil {
		t.Error("GetBook : Error: ", err)
	}

	if book.Book[0] != books.Books[0] {
		t.Error("GetBook: Names should match: ", book)
	}

	// chapters
	chapters, err := app.GetBookChapters(bookId)
	if err != nil {
		t.Error("GetBookChapters : Error: ", err)
	}

	if chapters.Total != 22 {
		t.Error("GetBookChapters: Wrong total ", chapters.Total)
	}
}

func TestBookOffset(t *testing.T) {
	app := lotr.New("")

	options := lotr.NewGetOptionsOffset(1)

	books, err := app.ListBooksOptions(options)
	if err != nil {
		t.Error("Error ListBooks ", err)
	}

	if books.Total != 3 {
		t.Error("Should have total 3 books ", books)
	}

	if len(books.Books) != 2 {
		t.Error("Should have returned 2 books with offset 1 ", books)
	}
}

func TestBookPage(t *testing.T) {
	app := lotr.New("")

	options := lotr.NewGetOptionsPageLimited(2, 2)

	books, err := app.ListBooksOptions(options)
	if err != nil {
		t.Error("Error ListBooks ", err)
	}

	if books.Total != 3 {
		t.Error("Should have total 3 books ", books)
	}

	if len(books.Books) != 1 {
		t.Error("Should have returned 1 books with page 2 limit 2 ", books)
	}
}

func TestOptions(t *testing.T) {
	app := lotr.New("")

	options := lotr.NewGetOptions()
	options.Offset = -1

	_, err := app.ListBooksOptions(options)
	if err == nil {
		t.Error("Should have thrown offset error ")
	}
}
