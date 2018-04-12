package filters

import (
	"minibbs/models"
	"regexp"
	"strconv"

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

var IsTopicUser = func(ctx *context.Context) {
	topicId, _ := strconv.Atoi(ctx.Input.Param(":id"))
	_, currUser := IsLogin(ctx)

	isAdmin := false

	currRoles := models.RoleManager.FindRolesByUser(&currUser)

	for _, v := range currRoles {
		if v.Name == models.ADMIN {
			isAdmin = true
		}
	}

	if !isAdmin {
		//whether is current user's topic
		topicUser := models.TopicManager.FindTopicById(topicId).User
		if topicUser.Username != currUser.Username {
			ctx.WriteString("你无权限访问此页面")
		}
	}
}
