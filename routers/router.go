package routers

import (
	"catapi/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.CatController{}, "get:Index")
	web.Router("/stream-breed", &controllers.CatController{}, "get:StreamBreed")
	web.Router("/random", &controllers.CatController{}, "get:GetRandomImages")
	web.Router("/vote", &controllers.CatController{}, "post:Vote")
	web.Router("/favorites", &controllers.CatController{}, "get:GetFavorites")
}
