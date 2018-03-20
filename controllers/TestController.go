package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) TestActive() {
	out := c.GetString("in")
	fmt.Println(out)
	c.Data["PageTitle"] = "用户注册"
	c.Layout = "layout/layout.tpl"
	c.TplName = "basic.tpl"
}
