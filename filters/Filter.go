package filters

import (
	"fmt"
	"minibbs/models"
	"regexp"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// IsLogin check user is login from cookie
func IsLogin(ctx *context.Context) (bool, models.User) {
	token, exsit := ctx.GetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"))
	user := models.User{}
	if exsit {
		exsit, user = models.UserManager.FindUserByToken(token)
	}
	return exsit, user
}

var HasPermission = func(ctx *context.Context) {
	ok, user := IsLogin(ctx)
	if !ok {
		ctx.Redirect(302, "/login")
	} else {
		permissions := models.UserManager.FindPermissionByUser(user.Id)
		fmt.Printf("permission{%v}\n", permissions)
		url := ctx.Request.RequestURI
		beego.Debug("url: ", url)
		var flag = false
		for _, v := range permissions {
			if a, _ := regexp.MatchString(v.Url, url); a {
				flag = true
				break
			}
		}
		if !flag {
			ctx.WriteString("你没有权限访问这个页面")
		}
	}
}

var FilterUser = func(ctx *context.Context) {
	ok, _ := IsLogin(ctx)
	if !ok {
		ctx.Redirect(302, "/login")
	}
}
