// catapi/models/cat_model.go
package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Define the structure for cat breed info
type CatInfo struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Id           string `json:"id"`
	WikipediaUrl string `json:"wikipedia_url"`
}

// Function to fetch cat info from external API
func FetchCatInfo(breedId string) (*CatInfo, error) {
	url := fmt.Sprintf("https://api.thecatapi.com/v1/breeds/%s", breedId)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var catInfo CatInfo
	if err := json.NewDecoder(resp.Body).Decode(&catInfo); err != nil {
		return nil, err
	}

	return &catInfo, nil
}
