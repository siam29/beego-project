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

// GetCatImages handles the API request to fetch multiple cat images by breed
func (c *CatController) GetCatImages() {
	apiKey, err := web.AppConfig.String("cat_api_key")
	if err != nil {
		c.Ctx.WriteString("Error retrieving API key: " + err.Error())
		return
	}

	// Get the breed from the query parameters
	breed := c.GetString("breed")
	if breed == "" {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		c.Ctx.WriteString("No breed specified")
		return
	}

	// URL to fetch multiple images (limit set to 5)
	url := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?breed_ids=%s&limit=5", breed)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.WriteString("Error creating request: " + err.Error())
		return
	}
	req.Header.Set("x-api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.WriteString("Error fetching cat images: " + err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.WriteString("Error reading response: " + err.Error())
		return
	}

	var images []struct {
		URL string `json:"url"`
	}
	if err := json.Unmarshal(body, &images); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Ctx.WriteString("Error parsing JSON: " + err.Error())
		return
	}

	// Prepare the JSON response with image URLs
	var imageURLs []string
	for _, img := range images {
		imageURLs = append(imageURLs, img.URL)
	}

	response := map[string]interface{}{
		"ImageURLs": imageURLs,
		"Breed":     breed,
	}

	// Send JSON response
	c.Data["json"] = response
	c.ServeJSON()
}
