package lotr

import (
	"encoding/json"
	"fmt"
	"log"
	"lotr/endPoints"
	"lotr/lotrResponse"
)

func (this *App) ListQuotesOptions(options *GetOptions) (*lotrResponse.ListQuotes, error) {
	log.Println("ListQuotesOptions Called")

	resp, err := this.ApiOptions(endPoints.LIST_QUOTES, options, nil)
	if err != nil {
		log.Println("ListQuotesOptions: Error: ", err)
		return nil, err
	}

	var result *lotrResponse.ListQuotes
	if err := json.Unmarshal(resp, &result); err != nil {
		err = fmt.Errorf("Cannot Unmarshal Json [%w]", err)
		log.Println("ListQuotesOptions: Error: ", err)
		return nil, err
	}

	log.Println("ListQuotesOptions: Done")
	return result, nil
}

func (this *App) ListQuotes() (*lotrResponse.ListQuotes, error) {
	return this.ListQuotesOptions(nil)
}

// same response get movie quote
// id - movie id
func (this *App) GetQuote(id string) (*lotrResponse.ListQuotes, error) {
	log.Println("GetQuote called", id)
	identifiers := map[string]string{
		"id": id,
	}

	resp, err := this.ApiOptions(endPoints.GET_QUOTE, nil, identifiers)
	if err != nil {
		log.Println("GetQuote: Error: ", err)
		return nil, err
	}

	var result *lotrResponse.ListQuotes
	if err := json.Unmarshal(resp, &result); err != nil { // Parse []byte to go struct pointer
		err = fmt.Errorf("Cannot Unmarshal Json : %s : [%w]", id, err)
		log.Println("GetQuote: Error: ", err)
		return nil, err
	}

	log.Println("GetQuote: Done: ", id)
	return result, nil
}
