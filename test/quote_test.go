package main

import (
	"lotr"
	"math"
	"testing"
)

const (
	TOTAL_QUOTES = 2390
)

func TestQuotes(t *testing.T) {
	app := lotr.New(ACCESS_CODE)

	quoteList, err := app.ListQuotes()
	if err != nil {
		t.Error("Error ListQuotes ", err)
	}

	if quoteList.Total != TOTAL_QUOTES {
		t.Errorf("Wrong Total Quotes: %d != %d (Should be)", quoteList.Total, TOTAL_QUOTES)
	}

	quoteId := quoteList.Quotes[0].Id
	quote, err := app.GetQuote(quoteId)
	if err != nil {
		t.Error("GetQuote : Error: ", err)
	}

	if quote.Quotes[0] != quoteList.Quotes[0] {
		t.Errorf("GetQuote: Quote should match: %v != %v", quote.Quotes[0], quoteList.Quotes[0])
	}

	// options
	limit := 100
	maxPage := int(math.Ceil(float64(TOTAL_QUOTES) / float64(limit)))
	options := lotr.NewGetOptionsPageLimited(maxPage, limit)
	quoteList, err = app.ListQuotesOptions(options)
	if err != nil {
		t.Error("Error ListQuotesOption ", err)
	}

	if quoteList.Total != TOTAL_QUOTES {
		t.Errorf("Wrong Total Quotes: %d != %d (Should be)", quoteList.Total, TOTAL_QUOTES)
	}

	if len(quoteList.Quotes) != (TOTAL_QUOTES % limit) {
		t.Errorf("Quotes Count for last page: %d != %d(Should be)", len(quoteList.Quotes), (TOTAL_QUOTES % limit))
	}
}
