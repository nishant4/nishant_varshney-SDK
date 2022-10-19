package lotr

import (
	"encoding/json"
	"fmt"
	"log"
	"lotr/endPoints"
	"lotr/lotrResponse"
)

func (this *App) ListChaptersOptions(options *GetOptions) (*lotrResponse.ListChapters, error) {
	log.Println("ListChaptersOptions Called")

	resp, err := this.ApiOptions(endPoints.LIST_CHAPTERS, options, nil)
	if err != nil {
		log.Println("ListChaptersOptions: Error: ", err)
		return nil, err
	}

	var result *lotrResponse.ListChapters
	if err := json.Unmarshal(resp, &result); err != nil {
		err = fmt.Errorf("Cannot Unmarshal Json [%w]", err)
		log.Println("ListChaptersOptions: Error: ", err)
		return nil, err
	}

	log.Println("ListChaptersOptions: Done")
	return result, nil
}

func (this *App) ListChapters() (*lotrResponse.ListChapters, error) {
	return this.ListChaptersOptions(nil)
}

// same response get movie quote
// id - movie id
func (this *App) GetChapter(id string) (*lotrResponse.ListChapters, error) {
	log.Println("GetChapter called", id)
	identifiers := map[string]string{
		"id": id,
	}

	resp, err := this.ApiOptions(endPoints.GET_CHAPTER, nil, identifiers)
	if err != nil {
		log.Println("GetChapter: Error: ", err)
		return nil, err
	}

	var result *lotrResponse.ListChapters
	if err := json.Unmarshal(resp, &result); err != nil { // Parse []byte to go struct pointer
		err = fmt.Errorf("Cannot Unmarshal Json : %s : [%w]", id, err)
		log.Println("GetChapter: Error: ", err)
		return nil, err
	}

	log.Println("GetChapter: Done: ", id)
	return result, nil
}
