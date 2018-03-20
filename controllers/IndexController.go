package controllers

import (
	"fmt"
	"minibbs/filters"
	"minibbs/models"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

// Index .
func (c *IndexController) Index() {
	c.Data["PageTitle"] = "首页"
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
	p, _ := strconv.Atoi(c.Ctx.Input.Query("p"))
	if p == 0 {
		p = 1
	}
	size, _ := beego.AppConfig.Int("page.size")
	s, _ := strconv.Atoi(c.Ctx.Input.Query("s"))
	c.Data["S"] = s
	section := models.Section{Id: s}
	c.Data["Page"] = models.TopicManager.PageTopic(p, size, &section)
	c.Data["Sections"] = models.FindAllSection()
	c.Layout = "layout/layout.tpl"
	c.TplName = "index.tpl"
}

// LoginPage .
func (c *IndexController) LoginPage() {
	IsLogin, _ := filters.IsLogin(c.Ctx)
	if IsLogin {
		c.Redirect("/", 302)
	} else {
		beego.ReadFromRequest(&c.Controller)
		u := models.UserManager.FindPermissionByUser(1)
		beego.Debug(u)
		c.Data["PageTitle"] = "登录"
		c.Layout = "layout/layout.tpl"
		c.TplName = "login.tpl"
	}
}

// Login .
func (c *IndexController) Login() {
	flash := beego.NewFlash()
	username, password := c.Input().Get("username"), c.Input().Get("password")
	if flag, user := models.UserManager.Login(username, password); flag {
		c.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), user.Token, 30*24*60*60, "/", beego.AppConfig.String("cookie.domain"), false, true)
		c.Redirect("/", 302)
	} else {
		flash.Error("用户名或密码错误")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
	}
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

	if exsit, _ := models.UserManager.FindUserByUserName(username); exsit {
		flash.Error("用户名已被注册")
		flash.Store(&c.Controller)
		c.Redirect("/register", http.StatusFound)
		return
	}

	// if exsit, _ := models.UserManager.FindUserByUserEmail(email); exsit {
	// 	flash.Error("邮箱已被注册")
	// 	flash.Store(&c.Controller)
	// 	c.Redirect("/register", http.StatusFound)
	// 	return
	// }

	authURL := models.EmailManager.GenerateAuthURL(email)
	models.EmailManager.SetTheme("用户帐号激活") //设置主题
	models.EmailManager.SetEmailContent(authURL)

	err := models.EmailManager.InitSendCfg(email, username)
	if err != nil {
		flash.Error("发送注册邮件初始化时发生错误，请联系管理员")
		flash.Store(&c.Controller)
		c.Redirect("/register", http.StatusFound)
		return
	}

	err = models.EmailManager.SendEmail()
	if err != nil {
		flash.Error("发送注册邮件时发生错误，请联系管理员")
		flash.Store(&c.Controller)
		c.Redirect("/register", http.StatusFound)
		return
	}

	user := models.User{
		Username: username,
		Password: password,
		Email:    email,
		Image:    "/static/imgs/default.png",
	}

	models.UserManager.SaveUser(&user)

	flash.Success("注册验证邮件已经发送到您的邮箱，请激活后再登录")
	flash.Store(&c.Controller)
	c.Redirect("/register", http.StatusFound)
	return
}

// ActiveAccount activation user account by check email
func (c *IndexController) ActiveAccount() {
	flash := beego.NewFlash()
	token := c.GetString("token")
	fmt.Println("token: " + token)

	if models.EmailManager != nil {
		isAccess, email := models.EmailManager.CheckEmailURL(token)

		if isAccess {
			err := models.UserManager.ActiveAccount(email)
			if err != nil {
				// glog.Errorf("active user by email error[%s]\n", err.Error())
				flash.Error("激活账户时发生错误，请联系管理员")
				return
			}
		}

		flash.Success("激活账户成功")
		return
	}

	flash.Error("发送注册邮件时发生错误，请联系管理员")
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
	c.Data["PageTitle"] = "关于"
	c.Layout = "layout/layout.tpl"
	c.TplName = "about.tpl"
}
