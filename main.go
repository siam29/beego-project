package main

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

func (c *CatController) Get() {
	// Get the breed parameter from the URL query
	breed := c.GetString("breed")
	if breed == "" {
		breed = "beng" // default to Bengal if no breed is selected
	}

	// Fetch images from TheCatAPI based on the breed
	url := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?breed_ids=%s&limit=5", breed)
	resp, err := http.Get(url)
	if err != nil {
		c.Ctx.WriteString("Error fetching images: " + err.Error())
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	// Parse the JSON response
	var images []struct {
		URL string `json:"url"`
	}
	err = json.Unmarshal(body, &images)
	if err != nil {
		c.Ctx.WriteString("Error parsing images: " + err.Error())
		return
	}

	// Pass the images and the selected breed to the template
	c.Data["Images"] = images
	c.Data["SelectedBreed"] = breed
	c.TplName = "cat_search.tpl"
}

func main() {
	web.Router("/", &CatController{})
	web.Run()
}
