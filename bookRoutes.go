package lotr

import (
	"encoding/json"
	"fmt"
	"log"
	"lotr/endPoints"
	"lotr/lotrResponse"
)

func (this *App) ListBooksOptions(options *GetOptions) (*lotrResponse.ListBooks, error) {
	log.Println("ListBooksOptions Called")

	resp, err := this.ApiOptions(endPoints.LIST_BOOKS, options, nil)
	if err != nil {
		log.Println("ListBooksOptions: Error: ", err)
		return nil, err
	}

	var result *lotrResponse.ListBooks
	if err := json.Unmarshal(resp, &result); err != nil {
		err = fmt.Errorf("Cannot Unmarshal Json [%w]", err)
		log.Println("ListBooksOptions: Error: ", err)
		return nil, err
	}

	log.Println("ListBooksOptions: Done")
	return result, nil
}

func (this *App) ListBooks() (*lotrResponse.ListBooks, error) {
	return this.ListBooksOptions(nil)
}

func (this *App) GetBook(id string) (*lotrResponse.ListBooks, error) {
	log.Println("GetBook called", id)
	identifiers := map[string]string{
		"id": id,
	}

	resp, err := this.ApiOptions(endPoints.GET_BOOK, nil, identifiers)
	if err != nil {
		log.Println("GetBook: Error: ", err)
		return nil, err
	}

	var result *lotrResponse.ListBooks
	if err := json.Unmarshal(resp, &result); err != nil { // Parse []byte to go struct pointer
		err = fmt.Errorf("Cannot Unmarshal Json : %s : [%w]", id, err)
		log.Println("GetBook: Error: ", err)
		return nil, err
	}

	log.Println("GetBook: Done: ", id)
	return result, nil
}

func (this *App) GetBookChaptersOptions(id string, options *GetOptions) (*lotrResponse.ListChapters, error) {
	log.Println("GetBookChaptersOptions called", id)
	identifiers := map[string]string{
		"id": id,
	}

	resp, err := this.ApiOptions(endPoints.GET_BOOK_CHAPTERS, options, identifiers)
	if err != nil {
		log.Println("GetBookChaptersOptions: Error: ", err)
		return nil, err
	}

	var result *lotrResponse.ListChapters
	if err := json.Unmarshal(resp, &result); err != nil { // Parse []byte to go struct pointer
		err = fmt.Errorf("Cannot Unmarshal Json : %s : [%w]", id, err)
		log.Println("GetBookChaptersOptions: Error: ", err)
		return nil, err
	}

	log.Println("GetBookChaptersOptions: Done: ", id)
	return result, nil
}

func (this *App) GetBookChapters(id string) (*lotrResponse.ListChapters, error) {
	return this.GetBookChaptersOptions(id, nil)
}
