package lotr

import (
	"encoding/json"
	"fmt"
	"log"
	"lotr/endPoints"
	"lotr/lotrResponse"
)

func (this *App) ListCharactersOptions(options *GetOptions) (*lotrResponse.ListCharacters, error) {
	log.Println("ListCharactersOptions Called")

	resp, err := this.ApiOptions(endPoints.LIST_CHARACTERS, options, nil)
	if err != nil {
		log.Println("ListCharactersOptions: Error: ", err)
		return nil, err
	}

	var result *lotrResponse.ListCharacters
	if err := json.Unmarshal(resp, &result); err != nil {
		err = fmt.Errorf("Cannot Unmarshal Json [%w]", err)
		log.Println("ListCharactersOptions: Error: ", err)
		return nil, err
	}

	log.Println("ListCharactersOptions: Done")
	return result, nil
}

func (this *App) ListCharacters() (*lotrResponse.ListCharacters, error) {
	return this.ListCharactersOptions(nil)
}
