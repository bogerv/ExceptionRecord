package controllers

import (
	"html/template"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.XSRFExpire = 7200
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "index.tpl"
}
