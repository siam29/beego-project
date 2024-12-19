package main

import (
    "github.com/beego/beego/v2/server/web"
    "catapi/controllers"
)

func main() {
    // Serve the root route with the template
    web.Router("/", &controllers.CatController{}, "get:GetSearchPage")
    // Handle fetching cat image by breed
    web.Router("/catimage", &controllers.CatController{}, "get:GetCatImage")
    
    // Run the server
    web.Run()
}
