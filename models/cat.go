package models

// CatImage holds the URL of the cat image and its breed
type CatImage struct {
	URL   string `json:"url"`
	Breed string `json:"breed"`
}
