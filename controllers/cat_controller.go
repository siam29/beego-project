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
	URL   string `json:"url"`
	Breed string `json:"breed"`
}

// GetSearchPage renders the cat search page
func (c *CatController) GetSearchPage() {
	// This will render the cat_search.tpl template when the user visits the root route
	c.TplName = "cat_search.tpl"
}

// GetCatImage handles the API request to fetch cat images by breed
func (c *CatController) GetCatImage() {
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

	// URL with breed filter
	url := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?breed_ids=%s", breed)

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
		c.Ctx.WriteString("Error fetching cat image: " + err.Error())
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

	// Prepare the JSON response
	response := map[string]interface{}{
		"ImageURL": "",
		"Breed":    breed,
	}

	if len(images) > 0 {
		response["ImageURL"] = images[0].URL
	}

	// Send JSON response
	c.Data["json"] = response
	c.ServeJSON()
}
