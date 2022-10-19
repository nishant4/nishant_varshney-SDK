package lotr

import (
	"encoding/json"
	"fmt"
	"log"
	"lotr/endPoints"
	"lotr/lotrResponse"
)

func (this *App) ListMoviesOptions(options *GetOptions) (*lotrResponse.ListMovies, error) {
	log.Println("ListMoviesOptions Called")

	resp, err := this.ApiOptions(endPoints.LIST_MOVIES, options, nil)
	if err != nil {
		log.Println("ListMoviesOptions: Error: ", err)
		return nil, err
	}

	var result *lotrResponse.ListMovies
	if err := json.Unmarshal(resp, &result); err != nil {
		err = fmt.Errorf("Cannot Unmarshal Json [%w]", err)
		log.Println("ListMoviesOptions: Error: ", err)
		return nil, err
	}

	log.Println("ListMoviesOptions: Done")
	return result, nil
}

func (this *App) ListMovies() (*lotrResponse.ListMovies, error) {
	return this.ListMoviesOptions(nil)
}

func (this *App) GetMovie(id string) (*lotrResponse.ListMovies, error) {
	log.Println("GetMovie called", id)
	identifiers := map[string]string{
		"id": id,
	}

	resp, err := this.ApiOptions(endPoints.GET_MOVIE, nil, identifiers)
	if err != nil {
		log.Println("GetMovie: Error: ", err)
		return nil, err
	}

	var result *lotrResponse.ListMovies
	if err := json.Unmarshal(resp, &result); err != nil { // Parse []byte to go struct pointer
		err = fmt.Errorf("Cannot Unmarshal Json : %s : [%w]", id, err)
		log.Println("GetMovie: Error: ", err)
		return nil, err
	}

	log.Println("GetMovie: Done: ", id)
	return result, nil
}

func (this *App) GetMovieQuotesOptions(id string, options *GetOptions) (*lotrResponse.ListQuotes, error) {
	log.Println("GetMovieQuotesOptions called", id)
	identifiers := map[string]string{
		"id": id,
	}

	resp, err := this.ApiOptions(endPoints.GET_MOVIE_QUOTES, options, identifiers)
	if err != nil {
		log.Println("GetMovieQuotesOptions: Error: ", err)
		return nil, err
	}

	var result *lotrResponse.ListQuotes
	if err := json.Unmarshal(resp, &result); err != nil { // Parse []byte to go struct pointer
		err = fmt.Errorf("Cannot Unmarshal Json : %s : [%w]", id, err)
		log.Println("GetMovieQuotesOptions: Error: ", err)
		return nil, err
	}

	log.Println("GetMovieQuotesOptions: Done: ", id)
	return result, nil
}

func (this *App) GetMovieQuotes(id string) (*lotrResponse.ListQuotes, error) {
	return this.GetMovieQuotesOptions(id, nil)
}
