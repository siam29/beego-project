package routers

import (
	"catapi/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/cat/getBreedData", &controllers.CatController{}, "get:GetBreedData")
}
