package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	//"catapi/models" // Make sure this path is correct based on your project structure
	"io"
	"log"

	"github.com/beego/beego/v2/server/web"
)
type CatController struct {
	web.Controller
}


type CatBreed struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Origin      string `json:"origin"`
	WikipediaURL string `json:"wikipedia_url"`
	 
}

type CatImage struct {
	ID     string     `json:"id"`
	URL    string     `json:"url"`
	Breeds []CatBreed `json:"breeds"`
}

var (
	catBreeds []CatBreed
	favorites map[string]bool
	mu        sync.RWMutex
)

func init() {
	favorites = make(map[string]bool)
	resp, err := http.Get("https://api.thecatapi.com/v1/breeds")
	if err != nil {
		log.Fatalf("Failed to fetch cat breeds: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	if err := json.Unmarshal(body, &catBreeds); err != nil {
		log.Fatalf("Failed to parse cat breeds: %v", err)
	}
}


// Handle the main page
func (c *CatController) Index() {
	mu.RLock()
	c.Data["Breeds"] = catBreeds
	mu.RUnlock()
	c.TplName = "index.html"
}

// Handle streaming breed data
// func (c *CatController) StreamBreed() {
// 	breedID := c.GetString("breed")
// 	if breedID == "" {
// 		breedID = "beng"
// 	}

// 	w := c.Ctx.ResponseWriter
// 	w.Header().Set("Content-Type", "text/event-stream")
// 	w.Header().Set("Cache-Control", "no-cache")
// 	w.Header().Set("Connection", "keep-alive")

// 	// Get breed description and name
// 	var description string
// 	var name string
// 	var origin string
// 	var wikipediaURL string
// 	mu.RLock()
// 	for _, breed := range catBreeds {
// 		if breed.ID == breedID {
// 			description = breed.Description
// 			name = breed.Name
// 			origin = breed.Origin
// 			wikipediaURL = breed.WikipediaURL
// 			break
// 		}
// 	}
// 	mu.RUnlock()

// 	// Send description event
// 	fmt.Fprintf(w, "event: description\ndata: %s\n\n", description)
// 	w.Flush()

// 	// Send name event
// 	fmt.Fprintf(w, "event: name\ndata: %s\n\n", name)
// 	w.Flush()

// 	fmt.Fprintf(w, "event: origin\ndata: %s\n\n", origin)
// 	w.Flush()

// 	// Send Wikipedia link event
// 	fmt.Fprintf(w, "event: wikipedia\ndata: %s\n\n", wikipediaURL)
// 	w.Flush()

// 	// Create channels for communication
// 	done := make(chan struct{})
// 	defer close(done)

// 	// Start image streaming
// 	go func() {
// 		ticker := time.NewTicker(3 * time.Second)
// 		defer ticker.Stop()
// 		count := 0

// 		for {
// 			select {
// 			case <-done:
// 				return
// 			case <-ticker.C:
// 				if count >= 5 { // Stop after 5 images
// 					return
// 				}

// 				resp, err := http.Get(fmt.Sprintf("https://api.thecatapi.com/v1/images/search?breed_ids=%s", breedID))
// 				if err != nil {
// 					continue
// 				}

// 				var images []CatImage
// 				json.NewDecoder(resp.Body).Decode(&images)
// 				resp.Body.Close()

// 				if len(images) > 0 {
// 					fmt.Fprintf(w, "event: image\ndata: %s\n\n", images[0].URL)
// 					w.Flush()
// 					count++
// 				}
// 			}
// 		}
// 	}()

// 	// Keep connection alive
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-c.Ctx.ResponseWriter.CloseNotify():
// 			return
// 		case <-ticker.C:
// 			fmt.Fprintf(w, ": keepalive\n\n")
// 			w.Flush()
// 		}
// 	}
// }


// GO Channel format

func (c *CatController) StreamBreed() {
	breedID := c.GetString("breed")
	if breedID == "" {
		breedID = "beng"
	}

	w := c.Ctx.ResponseWriter
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Channels for events and control
	descriptionChan := make(chan string)
	nameChan := make(chan string)
	originChan := make(chan string)
	wikipediaChan := make(chan string)
	imageChan := make(chan string)
	done := make(chan struct{})

	defer func() {
		close(descriptionChan)
		close(nameChan)
		close(originChan)
		close(wikipediaChan)
		close(imageChan)
		//close(done)
	}()

	// Fetch breed information and send to channels
	go func() {
		var description, name, origin, wikipediaURL string
		mu.RLock()
		for _, breed := range catBreeds {
			if breed.ID == breedID {
				description = breed.Description
				name = breed.Name
				origin = breed.Origin
				wikipediaURL = breed.WikipediaURL
				break
			}
		}
		mu.RUnlock()

		descriptionChan <- fmt.Sprintf("event: description\ndata: %s\n\n", description)
		nameChan <- fmt.Sprintf("event: name\ndata: %s\n\n", name)
		originChan <- fmt.Sprintf("event: origin\ndata: %s\n\n", origin)
		wikipediaChan <- fmt.Sprintf("event: wikipedia\ndata: %s\n\n", wikipediaURL)
	}()

	// Fetch images periodically and send to the image channel
	go func() {
		ticker := time.NewTicker(3 * time.Second)
		defer ticker.Stop()
		count := 0

		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				if count >= 5 { // Stop after 5 images
					return
				}

				resp, err := http.Get(fmt.Sprintf("https://api.thecatapi.com/v1/images/search?breed_ids=%s", breedID))
				if err != nil {
					continue
				}

				var images []CatImage
				if err := json.NewDecoder(resp.Body).Decode(&images); err == nil && len(images) > 0 {
					imageChan <- fmt.Sprintf("event: image\ndata: %s\n\n", images[0].URL)
					count++
				}
				resp.Body.Close()
			}
		}
	}()

	// Keep connection alive with periodic keep-alive messages
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				fmt.Fprintf(w, ": keepalive\n\n")
				w.Flush()
			}
		}
	}()

	// Main loop to send events to the client
	for {
		select {
		case msg := <-descriptionChan:
			fmt.Fprint(w, msg)
			w.Flush()
		case msg := <-nameChan:
			fmt.Fprint(w, msg)
			w.Flush()
		case msg := <-originChan:
			fmt.Fprint(w, msg)
			w.Flush()
		case msg := <-wikipediaChan:
			fmt.Fprint(w, msg)
			w.Flush()
		case msg := <-imageChan:
			fmt.Fprint(w, msg)
			w.Flush()
		case <-c.Ctx.ResponseWriter.CloseNotify():
			close(done)
			return
		}
	}
}



// Handle random images for voting
// func (c *CatController) GetRandomImages() {
// 	resp, err := http.Get("https://api.thecatapi.com/v1/images/search?limit=10")
// 	if err != nil {
// 		c.Data["json"] = map[string]string{"error": err.Error()}
// 		c.ServeJSON()
// 		return
// 	}
// 	defer resp.Body.Close()

// 	var images []CatImage
// 	if err := json.NewDecoder(resp.Body).Decode(&images); err != nil {
// 		c.Data["json"] = map[string]string{"error": err.Error()}
// 		c.ServeJSON()
// 		return
// 	}

// 	c.Data["json"] = images
// 	c.ServeJSON()
// }

// GO Channel format

func (c *CatController) GetRandomImages() {
	imageChan := make(chan []CatImage) // Channel for images
	errorChan := make(chan error)     // Channel for errors

	// Goroutine to fetch random images
	go func() {
		resp, err := http.Get("https://api.thecatapi.com/v1/images/search?limit=10")
		if err != nil {
			errorChan <- err
			close(imageChan)
			close(errorChan)
			return
		}
		defer resp.Body.Close()

		var images []CatImage
		if err := json.NewDecoder(resp.Body).Decode(&images); err != nil {
			errorChan <- err
			close(imageChan)
			close(errorChan)
			return
		}

		imageChan <- images
		close(imageChan)
		close(errorChan)
	}()

	// Listen on channels for the result
	select {
	case images := <-imageChan:
		if images != nil {
			c.Data["json"] = images
		} else {
			c.Data["json"] = map[string]string{"error": "No images fetched"}
		}
	case err := <-errorChan:
		if err != nil {
			c.Data["json"] = map[string]string{"error": err.Error()}
		}
	}

	c.ServeJSON()
}


// Handle voting/favorites
func (c *CatController) Vote() {
	var data struct {
		ImageID  string `json:"image_id"`
		VoteType string `json:"vote_type"`
	}

	if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&data); err != nil {
		c.Data["json"] = map[string]string{"error": "invalid request"}
		c.ServeJSON()
		return
	}

	if data.VoteType == "favorite" {
		mu.Lock()
		favorites[data.ImageID] = true
		mu.Unlock()
	}

	c.Data["json"] = map[string]string{"status": "success"}
	c.ServeJSON()
}

// Get favorites
// func (c *CatController) GetFavorites() {
// 	mu.RLock()
// 	favoriteIDs := make([]string, 0, len(favorites))
// 	for id := range favorites {
// 		favoriteIDs = append(favoriteIDs, id)
// 	}
// 	mu.RUnlock()

// 	var favoriteImages []CatImage
// 	for _, id := range favoriteIDs {
// 		resp, err := http.Get(fmt.Sprintf("https://api.thecatapi.com/v1/images/%s", id))
// 		if err != nil {
// 			continue
// 		}
// 		var img CatImage
// 		json.NewDecoder(resp.Body).Decode(&img)
// 		resp.Body.Close()
// 		favoriteImages = append(favoriteImages, img)
// 	}

// 	c.Data["json"] = favoriteImages
// 	c.ServeJSON()
// }

// GO Channel format

func (c *CatController) GetFavorites() {
	mu.RLock()
	favoriteIDs := make([]string, 0, len(favorites))
	for id := range favorites {
		favoriteIDs = append(favoriteIDs, id)
	}
	mu.RUnlock()

	// Create channels for fetching results
	imageChan := make(chan CatImage)   // Channel to send individual images
	errorChan := make(chan error)      // Channel to send errors
	doneChan := make(chan struct{})    // Channel to signal completion of all fetches

	// Goroutine to fetch all favorite images concurrently
	go func() {
		var favoriteImages []CatImage
		var fetchError error

		// Loop through each favorite ID and fetch images concurrently
		for _, id := range favoriteIDs {
			go func(id string) {
				resp, err := http.Get(fmt.Sprintf("https://api.thecatapi.com/v1/images/%s", id))
				if err != nil {
					errorChan <- err
					return
				}
				defer resp.Body.Close()

				var img CatImage
				if err := json.NewDecoder(resp.Body).Decode(&img); err != nil {
					errorChan <- err
					return
				}

				imageChan <- img
			}(id)
		}

		// Wait for all goroutines to finish and collect results
		for range favoriteIDs {
			select {
			case img := <-imageChan:
				favoriteImages = append(favoriteImages, img)
			case err := <-errorChan:
				// If there's an error, we can log it or handle it
				fetchError = err
			}
		}

		// Notify completion
		doneChan <- struct{}{}

		// Send final images or error
		if fetchError != nil {
			c.Data["json"] = map[string]string{"error": fetchError.Error()}
		} else {
			c.Data["json"] = favoriteImages
		}
	}()

	// Wait for the completion signal
	<-doneChan
	c.ServeJSON()
}
