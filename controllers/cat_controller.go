package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type CatController struct {
	web.Controller
}

type CatImage struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

// Function to fetch cat images from TheCatAPI
func (c *CatController) GetCatImages() {
	breed := c.GetString("breed")
	if breed == "" {
		c.Data["json"] = map[string]string{"message": "Breed not specified"}
		c.ServeJSON()
		return
	}

	// URL for TheCatAPI to get multiple images of the specified breed
	url := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?breed_ids=%s&limit=5", breed)

	// Make the HTTP request to TheCatAPI
	resp, err := http.Get(url)
	if err != nil {
		c.Data["json"] = map[string]string{"message": "Error fetching cat images"}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Data["json"] = map[string]string{"message": "Error reading response"}
		c.ServeJSON()
		return
	}

	// Unmarshal the JSON response into CatImage structure
	var images []CatImage
	if err := json.Unmarshal(body, &images); err != nil {
		c.Data["json"] = map[string]string{"message": "Error parsing response"}
		c.ServeJSON()
		return
	}

	// Return the list of image URLs as JSON
	var imageURLs []string
	for _, image := range images {
		imageURLs = append(imageURLs, image.URL)
	}

	c.Data["json"] = map[string]interface{}{
		"Images": imageURLs,
	}
	c.ServeJSON()
}

// Display the search page with a dropdown for selecting a breed
func (c *CatController) Get() {
	c.TplName = "cat_search.tpl"
}
