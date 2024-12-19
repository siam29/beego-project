package routers

import (
	"catapi/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	// Register the route for CatController
	web.Router("/", &controllers.CatController{})
	web.Router("/catimages", &controllers.CatController{}, "get:GetCatImages")
}
