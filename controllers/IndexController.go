package controllers

import (
	"minibbs/filters"
	"minibbs/models"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/sluu99/uuid"
)

type IndexController struct {
	beego.Controller
}

// Index .
func (c *IndexController) Index() {
	c.Data["PageTitle"] = "首页"
	isLogin, user := filters.IsLogin(c.Controller.Ctx)
	c.Data["IsLogin"] = isLogin
	c.Data["UserInfo"] = user

	page, _ := strconv.Atoi(c.Ctx.Input.Query("p"))
	if page == 0 {
		page = 1
	}
	size, _ := beego.AppConfig.Int("page.size")
	tagId, _ := strconv.Atoi(c.Ctx.Input.Query("tagId"))
	c.Data["TagId"] = tagId
	tag := models.Tag{Id: tagId}
	c.Data["Page"] = models.TopicManager.PageTopic(page, size, &tag)
	c.Data["Tags"] = models.TagManager.FindAllTag()
	c.Layout = "layout/layout.tpl"
	c.TplName = "index.tpl"
}

// LoginPage .
func (c *IndexController) LoginPage() {
	IsLogin, _ := filters.IsLogin(c.Ctx)
	beego.ReadFromRequest(&c.Controller) // for flash data
	if IsLogin {
		c.Redirect("/", 302)
	} else {
		c.Data["PageTitle"] = "登录"
		c.Layout = "layout/layout.tpl"
		c.TplName = "login.tpl"
	}
}

// Login .
func (c *IndexController) Login() {
	flash := beego.NewFlash()
	username := c.Input().Get("username")
	password := c.Input().Get("password")
	roleStr := c.Input().Get("role")

	hasRole := false

	exsit, user, err := models.UserManager.Login(username, password)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	roles := models.RoleManager.FindRolesByUser(user)
	for _, role := range roles {
		if roleStr == role.Name {
			hasRole = true
		}
	}

	if !hasRole {
		flash.Error("登录身份类型错误")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	if exsit {
		c.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), user.Token, 30*24*60*60, "/", beego.AppConfig.String("cookie.domain"), false, true)
		c.Redirect("/", 302)
		return
	}

	flash.Error("用户名或密码错误")
	flash.Store(&c.Controller)
	c.Redirect("/login", 302)
	return
}

// RegisterPage .
func (c *IndexController) RegisterPage() {
	isLogin, _ := filters.IsLogin(c.Ctx)

	if isLogin {
		c.Redirect("/", http.StatusFound)
		return
	}

	beego.ReadFromRequest(&c.Controller)
	c.Data["PageTitle"] = "用户注册"
	c.Layout = "layout/layout.tpl"
	c.TplName = "register.tpl"
	return

}

// Register .
func (c *IndexController) Register() {
	flash := beego.NewFlash()
	username, password, email := c.Input().Get("username"), c.Input().Get("password"), c.Input().Get("email")
	if len(username) == 0 || len(password) == 0 || len(email) == 0 {
		flash.Error("输入框不能为空")
		flash.Store(&c.Controller)
		c.Redirect("/register", http.StatusFound)
		return
	}

	var token = uuid.Rand().Hex() // user token

	user := models.User{
		Username: username,
		Password: password,
		Email:    email,
		Token:    token,
		Image:    "/static/imgs/default.png",
		Active:   true,
	}

	if exsit, _ := models.UserManager.FindUserByUserName(username); exsit {
		flash.Error("用户名已被注册")
		flash.Store(&c.Controller)
		c.Redirect("/register", http.StatusFound)
		return
	}

	if exsit, _ := models.UserManager.FindUserByUserEmail(email); exsit {
		flash.Error("邮箱已被注册")
		flash.Store(&c.Controller)
		c.Redirect("/register", http.StatusFound)
		return
	}

	if err := models.UserManager.SaveUser(&user); err != nil {
		flash.Error("注册用户失败:" + err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/register", http.StatusFound)
		return
	}

	// 普通用户赋值
	role := models.RoleManager.FindRoleByName("普通用户")
	models.UserManager.SaveUserRole(user.Id, role.Id)

	flash.Success("注册成功")
	flash.Store(&c.Controller)
	c.Redirect("/register", http.StatusFound)
	return
}

// Logout .
func (c *IndexController) Logout() {
	c.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), "", -1, "/", beego.AppConfig.String("cookie.domain"), false, true)
	c.Redirect("/", 302)
}

// About .
func (c *IndexController) About() {
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
	c.Data["PageTitle"] = "公告"
	c.Layout = "layout/layout.tpl"
	c.TplName = "about.tpl"
}
