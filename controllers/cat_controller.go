package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/server/web"
)

type CatController struct {
	web.Controller
}

// BreedData struct to hold breed description and images
type BreedData struct {
	Description string   `json:"description"`
	Images      []string `json:"images"`
}

// Get handles the request for cat breeds and their data
func (c *CatController) Get() {
	// Get breed from the URL parameter
	breedID := c.GetString("breed")

	// Fetch breed data
	breedData, err := c.fetchBreedData(breedID)
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	// Send breed data to the frontend (as JSON)
	c.Data["json"] = breedData
	c.ServeJSON()
}

// fetchBreedData fetches description and images for the selected breed
func (c *CatController) fetchBreedData(breedID string) (BreedData, error) {
	// Access the configuration values from app.conf
	apiKey := config.String("api_key")
	baseURL := config.String("base_url")

	// Fetch breed description
	descriptionURL := fmt.Sprintf("%sbreeds", baseURL)
	req, err := http.NewRequest("GET", descriptionURL, nil)
	if err != nil {
		return BreedData{}, fmt.Errorf("Error creating request for breed data: %v", err)
	}
	req.Header.Set("x-api-key", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return BreedData{}, fmt.Errorf("Error fetching breed data: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return BreedData{}, fmt.Errorf("Error reading breed response body: %v", err)
	}

	var breeds []map[string]interface{}
	err = json.Unmarshal(body, &breeds)
	if err != nil {
		return BreedData{}, fmt.Errorf("Error unmarshalling breed data: %v", err)
	}

	var description string
	for _, breed := range breeds {
		if breed["id"] == breedID {
			description = breed["description"].(string)
			break
		}
	}

	// Fetch images of the selected breed
	imagesURL := fmt.Sprintf("%simages/search?breed_ids=%s&limit=5", baseURL, breedID)
	req, err = http.NewRequest("GET", imagesURL, nil)
	if err != nil {
		return BreedData{}, fmt.Errorf("Error creating request for cat images: %v", err)
	}
	req.Header.Set("x-api-key", apiKey)

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return BreedData{}, fmt.Errorf("Error fetching cat images: %v", err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return BreedData{}, fmt.Errorf("Error reading images response body: %v", err)
	}

	var images []map[string]interface{}
	err = json.Unmarshal(body, &images)
	if err != nil {
		return BreedData{}, fmt.Errorf("Error unmarshalling image data: %v", err)
	}

	var imageURLs []string
	for _, image := range images {
		imageURLs = append(imageURLs, image["url"].(string))
	}

	return BreedData{
		Description: description,
		Images:      imageURLs,
	}, nil
}
