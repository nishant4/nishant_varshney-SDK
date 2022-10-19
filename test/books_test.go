package main

import (
	"lotr"
	"testing"
)

const (
	BOOK_ID             = "5cf5805fb53e011a64671582"
	TOTAL_BOOK_COUNT    = 3
	TOTAL_BOOK_CHAPTERS = 22
)

func TestBooks(t *testing.T) {
	app := lotr.New("")

	bookList, err := app.ListBooks()
	if err != nil {
		t.Fatal("Error ListBooks ", err)
	}

	if bookList.Total != TOTAL_BOOK_COUNT {
		t.Errorf("Wrong book count %d != %d (Should be)", bookList.Total, TOTAL_BOOK_COUNT)
	}

	bookId := bookList.Books[0].Id
	book, err := app.GetBook(bookId)
	if err != nil {
		t.Fatal("GetBook : Error: ", err)
	}

	if book.Books[0] != bookList.Books[0] {
		t.Errorf("GetBook: Books should match: %v != %v ", book.Books[0], bookList.Books[0])
	}

}

func TestGetChapters(t *testing.T) {
	app := lotr.New("")

	// chapters
	chapterList, err := app.GetBookChapters(BOOK_ID)
	if err != nil {
		t.Fatal("GetBookChapters : Error: ", err)
	}

	if chapterList.Total != TOTAL_BOOK_CHAPTERS {
		t.Error("GetBookChapters: Wrong total ", chapterList.Total)
	}

	// Chapters options
	offset := 10
	chapterList, err = app.GetBookChaptersOptions(BOOK_ID, lotr.NewGetOptionsOffset(offset))
	if err != nil {
		t.Fatal("GetBookChapters : Error: ", err)
	}

	if chapterList.Total != TOTAL_BOOK_CHAPTERS {
		t.Error("GetBookChapters: Wrong total ", chapterList.Total)
	}

	if len(chapterList.Chapters) != TOTAL_BOOK_CHAPTERS-offset {
		t.Error("GetBookChapters: Wrong total ", chapterList.Total)
	}
}

func TestBookOffset(t *testing.T) {
	app := lotr.New("")

	options := lotr.NewGetOptionsOffset(1)

	bookList, err := app.ListBooksOptions(options)
	if err != nil {
		t.Fatal("Error ListBooks ", err)
	}

	if bookList.Total != 3 {
		t.Error("Should have total 3 books ", bookList)
	}

	if len(bookList.Books) != 2 {
		t.Error("Should have returned 2 books with offset 1 ", bookList)
	}
}

func TestBookPage(t *testing.T) {
	app := lotr.New("")

	options := lotr.NewGetOptionsPageLimited(2, 2)

	bookList, err := app.ListBooksOptions(options)
	if err != nil {
		t.Fatal("Error ListBooks ", err)
	}

	if bookList.Total != 3 {
		t.Error("Should have total 3 books ", bookList)
	}

	if len(bookList.Books) != 1 {
		t.Error("Should have returned 1 books with page 2 limit 2 ", bookList)
	}
}

func TestOptions(t *testing.T) {
	app := lotr.New("")

	options := lotr.NewGetOptions()
	options.Offset = -1

	_, err := app.ListBooksOptions(options)
	if err == nil {
		t.Fatal("Should have thrown offset error ")
	}
}
