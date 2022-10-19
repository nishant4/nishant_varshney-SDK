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

func (this *App) GetBook(id string) (*lotrResponse.GetBook, error) {
	log.Println("GetBook called", id)
	identifiers := map[string]string{
		"id": id,
	}

	resp, err := this.ApiOptions(endPoints.GET_BOOK, nil, identifiers)
	if err != nil {
		log.Println("GetBook: Error: ", err)
		return nil, err
	}

	var result *lotrResponse.GetBook
	if err := json.Unmarshal(resp, &result); err != nil { // Parse []byte to go struct pointer
		err = fmt.Errorf("Cannot Unmarshal Json : %s : [%w]", id, err)
		log.Println("GetBook: Error: ", err)
		return nil, err
	}

	log.Println("GetBook: Done: ", id)
	return result, nil
}

func (this *App) GetBookChapters(id string) (*lotrResponse.GetBookChapters, error) {
	log.Println("GetBookChapters called", id)
	identifiers := map[string]string{
		"id": id,
	}

	resp, err := this.ApiOptions(endPoints.GET_BOOK_CHAPTERS, nil, identifiers)
	if err != nil {
		log.Println("GetBookChapters: Error: ", err)
		return nil, err
	}

	var result *lotrResponse.GetBookChapters
	if err := json.Unmarshal(resp, &result); err != nil { // Parse []byte to go struct pointer
		err = fmt.Errorf("Cannot Unmarshal Json : %s : [%w]", id, err)
		log.Println("GetBookChapters: Error: ", err)
		return nil, err
	}

	log.Println("GetBookChapters: Done: ", id)
	return result, nil
}

// func (this *App) GetBookChapters(id string) (*lotrResponse.GetBookChapters, error) {
// 	log.Println("GetBookChapters called", id)
// 	params := map[string]string{
// 		"id": id,
// 	}

// 	resp, err := network.HttpCallWithIdentifiers(this.conf.Route, this.conf.GetEndPoint(endPoints.GET_BOOK_CHAPTERS), params)
// 	if err != nil {
// 		log.Println("GetBookChapters: Error: ", id, err)
// 		return nil, err
// 	}

// 	var result *lotrResponse.GetBookChapters
// 	if err := json.Unmarshal(resp, &result); err != nil { // Parse []byte to go struct pointer
// 		errMsg := fmt.Errorf("GetBookChapters: Cannot Unmarshal Json : %s : [%w]", id, err)
// 		log.Println(errMsg)
// 		return nil, errMsg
// 	}
// 	log.Println("GetBookChapters: Done: ", id)
// 	return result, err
// }
