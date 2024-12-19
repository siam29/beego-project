package routers

import (
	"catapi/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	// Root route to handle index
	web.Router("/", &controllers.CatController{})

	// Route to handle fetching cat images by breed
	web.Router("/catimage", &controllers.CatController{}, "get:GetCatImage")
}
