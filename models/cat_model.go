package models

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
)

type CatBreed struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Origin      string `json:"origin"`
}

type CatImage struct {
	ID     string     `json:"id"`
	URL    string     `json:"url"`
	Breeds []CatBreed `json:"breeds"`
}

var (
	CatBreeds []CatBreed
	Favorites map[string]bool
	mu        sync.RWMutex
)

func InitCatBreeds() {
	Favorites = make(map[string]bool)

	resp, err := http.Get("https://api.thecatapi.com/v1/breeds")
	if err != nil {
		log.Fatalf("Failed to fetch cat breeds: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	if err := json.Unmarshal(body, &CatBreeds); err != nil {
		log.Fatalf("Failed to parse cat breeds: %v", err)
	}
}

func GetBreedDescription(breedID string) string {
	mu.RLock()
	defer mu.RUnlock()

	for _, breed := range CatBreeds {
		if breed.ID == breedID {
			return breed.Description
		}
	}
	return ""
}

func AddFavorite(imageID string) {
	mu.Lock()
	defer mu.Unlock()
	Favorites[imageID] = true
}

func GetFavoriteImages() []string {
	mu.RLock()
	defer mu.RUnlock()

	favoriteIDs := make([]string, 0, len(Favorites))
	for id := range Favorites {
		favoriteIDs = append(favoriteIDs, id)
	}
	return favoriteIDs
}
