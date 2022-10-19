package main

import (
	"lotr"
	"testing"
)

const (
	TOTAL_MOVIES       = 8
	MOVIE_ID           = "5cd95395de30eff6ebccde5d" // the return of the king
	TOTAL_MOVIE_QUOTES = 873
)

func TestMovies(t *testing.T) {
	app := lotr.New(ACCESS_CODE)

	movieList, err := app.ListMovies()
	if err != nil {
		t.Error("Error ListMovies ", err)
	}

	if movieList.Total != TOTAL_MOVIES {
		t.Errorf("Wrong Total Characters: %d != %d (Should be)", movieList.Total, TOTAL_MOVIES)
	}

	movieId := movieList.Movies[0].Id
	movie, err := app.GetMovie(movieId)
	if err != nil {
		t.Error("GetMovie : Error: ", err)
	}

	if movie.Movies[0] != movieList.Movies[0] {
		t.Errorf("GetMovie: Characters should match: %v != %v", movie.Movies[0], movieList.Movies[0])
	}

	// options
	offset := 3
	options := lotr.NewGetOptionsOffset(offset)
	movieList, err = app.ListMoviesOptions(options)
	if err != nil {
		t.Error("Error ListMovies ", err)
	}

	if movieList.Total != TOTAL_MOVIES {
		t.Errorf("Wrong Total Characters: %d != %d (Should be)", movieList.Total, TOTAL_MOVIES)
	}

	if len(movieList.Movies) != TOTAL_MOVIES-offset {
		t.Errorf("Movies Count with offset: %d != %d(Should be)", len(movieList.Movies), TOTAL_MOVIES)
	}
}

func TestMovieQuotes(t *testing.T) {
	app := lotr.New(ACCESS_CODE)

	movieQuotes, err := app.GetMovieQuotes(MOVIE_ID)

	if err != nil {
		t.Error("Get Error: ", err)
	}

	if movieQuotes.Total != TOTAL_MOVIE_QUOTES {
		t.Errorf("Movie Quote Count: %d != %d(Should be)", movieQuotes.Total, TOTAL_MOVIE_QUOTES)
	}

	offset := 10
	movieQuotes, err = app.GetMovieQuotesOptions(MOVIE_ID, lotr.NewGetOptionsOffset(offset))

	if err != nil {
		t.Error("Get Error: ", err)
	}

	if movieQuotes.Total != TOTAL_MOVIE_QUOTES {
		t.Errorf("Movie Quote Count: %d != %d (Should be)", movieQuotes.Total, TOTAL_MOVIE_QUOTES)
	}

	if len(movieQuotes.Quotes) != TOTAL_MOVIE_QUOTES-offset {
		t.Errorf("Movie Quote Count with offset: %d != %d (Should be)", movieQuotes.Total, TOTAL_MOVIE_QUOTES)
	}
}
