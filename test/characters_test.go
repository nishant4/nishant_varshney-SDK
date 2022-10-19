package main

import (
	"lotr"
	"testing"
)

const (
	TOTAL_CHARACTERS     = 933
	GANDALF_ID           = "5cd99d4bde30eff6ebccfea0"
	GANDALF_TOTAL_QUOTES = 216
)

func TestCharacters(t *testing.T) {
	app := lotr.New(ACCESS_CODE)

	charList, err := app.ListCharacters()
	if err != nil {
		t.Fatal("Error ListCharacters ", err)
	}

	if charList.Total != TOTAL_CHARACTERS {
		t.Errorf("Wrong Total Characters: %d: Should be: %d", charList.Total, TOTAL_CHARACTERS)
	}

	charId := charList.Characters[0].Id
	char, err := app.GetCharacter(charId)
	if err != nil {
		t.Fatal("GetCharacter : Error: ", err)
	}

	if char.Characters[0] != charList.Characters[0] {
		t.Errorf("GetCharacter: Characters should match: %v != %v", char.Characters[0], charList.Characters[0])
	}
}

func TestCharacterQuotes(t *testing.T) {
	app := lotr.New(ACCESS_CODE)

	charQuotes, err := app.GetCharacterQuotes(GANDALF_ID)

	if err != nil {
		t.Fatal("Get Error: ", err)
	}

	if charQuotes.Total != GANDALF_TOTAL_QUOTES {
		t.Errorf("Gandalf Quote Count: %d != %d(Should be)", charQuotes.Total, GANDALF_TOTAL_QUOTES)
	}

	offset := 10
	charQuotes, err = app.GetCharacterQuotesOptions(GANDALF_ID, lotr.NewGetOptionsOffset(offset))

	if err != nil {
		t.Fatal("Get Error: ", err)
	}

	if charQuotes.Total != GANDALF_TOTAL_QUOTES {
		t.Errorf("Gandalf Quote Count: %d != %d(Should be)", charQuotes.Total, GANDALF_TOTAL_QUOTES)
	}

	if len(charQuotes.Quotes) != GANDALF_TOTAL_QUOTES-offset {
		t.Errorf("Gandalf Quote Count with offset: %d != %d(Should be)", charQuotes.Total, GANDALF_TOTAL_QUOTES)
	}
}
