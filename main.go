package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/beego/beego/v2/server/web"
)

type CatController struct {
	web.Controller
}

type CatBreed struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CatImage struct {
	URL string `json:"url"`
}

var (
	catBreeds []CatBreed
)

// Fetch all breeds and store them in memory
func init() {
	resp, err := http.Get("https://api.thecatapi.com/v1/breeds")
	if err != nil {
		log.Fatalf("Failed to fetch cat breeds: %v", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &catBreeds); err != nil {
		log.Fatalf("Failed to parse cat breeds: %v", err)
	}
}

// Handle the initial cat search page
func (c *CatController) Search() {
	c.Data["Breeds"] = catBreeds
	c.TplName = "cat_search.tpl"
}

// Handle streaming cat data (images cycling and description)
func (c *CatController) StreamData() {
	breedID := c.GetString("breed")
	if breedID == "" {
		breedID = "beng" // Default to Bengal
	}

	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/event-stream")
	c.Ctx.ResponseWriter.Header().Set("Cache-Control", "no-cache")

	// Goroutine for sending updates
	go func() {
		// Fetch the breed description
		description := "Description not available."
		for _, breed := range catBreeds {
			if breed.ID == breedID {
				description = breed.Description
				break
			}
		}

		// Send the description once
		fmt.Fprintf(c.Ctx.ResponseWriter, "event: description\n")
		fmt.Fprintf(c.Ctx.ResponseWriter, "data: %s\n\n", description)
		c.Ctx.ResponseWriter.Flush()

		// Cycle through images every 3000ms
		for {
			resp, err := http.Get(fmt.Sprintf("https://api.thecatapi.com/v1/images/search?limit=1&breed_ids=%s", breedID))
			if err == nil {
				defer resp.Body.Close()
				body, _ := ioutil.ReadAll(resp.Body)
				var images []CatImage
				json.Unmarshal(body, &images)
				if len(images) > 0 {
					// Send image URL
					fmt.Fprintf(c.Ctx.ResponseWriter, "event: image\n")
					fmt.Fprintf(c.Ctx.ResponseWriter, "data: %s\n\n", images[0].URL)
					c.Ctx.ResponseWriter.Flush()
				}
			}
			time.Sleep(3 * time.Second)
		}
	}()

	// Keep the connection alive
	for {
		time.Sleep(1 * time.Second)
	}
}

func main() {
	web.Router("/", &CatController{}, "get:Search")
	web.Router("/stream", &CatController{}, "get:StreamData")
	web.Run()
}
