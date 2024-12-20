// catapi/controllers/main_controller.go
package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (c *MainController) Get() {
	c.TplName = "cat.html"
}
