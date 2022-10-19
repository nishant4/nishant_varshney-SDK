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

func (this *App) GetCharacter(id string) (*lotrResponse.ListCharacters, error) {
	log.Println("GetCharacter called", id)
	identifiers := map[string]string{
		"id": id,
	}

	resp, err := this.ApiOptions(endPoints.GET_CHARACTER, nil, identifiers)
	if err != nil {
		log.Println("GetCharacter: Error: ", err)
		return nil, err
	}

	var result *lotrResponse.ListCharacters
	if err := json.Unmarshal(resp, &result); err != nil { // Parse []byte to go struct pointer
		err = fmt.Errorf("Cannot Unmarshal Json : %s : [%w]", id, err)
		log.Println("GetCharacter: Error: ", err)
		return nil, err
	}

	log.Println("GetCharacter: Done: ", id)
	return result, nil
}

func (this *App) GetCharacterQuotesOptions(id string, options *GetOptions) (*lotrResponse.ListQuotes, error) {
	log.Println("GetCharacterQuotesOptions called", id)
	identifiers := map[string]string{
		"id": id,
	}

	resp, err := this.ApiOptions(endPoints.GET_CHARACTER_QUOTES, options, identifiers)
	if err != nil {
		log.Println("GetCharacterQuotesOptions: Error: ", err)
		return nil, err
	}

	var result *lotrResponse.ListQuotes
	if err := json.Unmarshal(resp, &result); err != nil { // Parse []byte to go struct pointer
		err = fmt.Errorf("Cannot Unmarshal Json : %s : [%w]", id, err)
		log.Println("GetCharacterQuotesOptions: Error: ", err)
		return nil, err
	}

	log.Println("GetCharacterQuotesOptions: Done: ", id)
	return result, nil
}

func (this *App) GetCharacterQuotes(id string) (*lotrResponse.ListQuotes, error) {
	return this.GetCharacterQuotesOptions(id, nil)
}
