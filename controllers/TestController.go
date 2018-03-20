package controllers

import (
	"minibbs/models"

	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) TestActive() {
	email := c.GetString("email")
	err := models.UserManager.ActiveAccount(email)
	if err != nil {
		// glog.Errorf("active user by email error[%s]\n", err.Error())
		c.Ctx.WriteString("激活账户时发生错误，请联系管理员 " + err.Error())
		return
	}

	c.Ctx.WriteString("激活成功")
}
