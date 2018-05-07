package controllers

import (
	"fmt"
	"minibbs/filters"
	"minibbs/models"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Detail() {
	username := c.Ctx.Input.Param(":username")
	ok, user := models.UserManager.FindUserByUserName(username)
	if ok {
		c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
		c.Data["PageTitle"] = "个人主页"
		c.Data["CurrentUserInfo"] = user
		c.Data["Topics"] = models.TopicManager.FindTopicByUser(&user, 5, 0, 0)
		c.Data["Replies"] = models.ReplyManager.FindReplyByUser(&user, 5, 0, 0)
	}
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/detail.tpl"
}

func (c *UserController) ToSetting() {
	beego.ReadFromRequest(&c.Controller)
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
	c.Data["PageTitle"] = "用户设置"
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/setting.tpl"
}

func (c *UserController) Setting() {
	flash := beego.NewFlash()
	email, signature := c.Input().Get("email"), c.Input().Get("signature")
	if len(email) > 0 {
		ok, _ := regexp.MatchString("^([a-z0-9A-Z]+[-|_|\\.]?)+[a-z0-9A-Z]@([a-z0-9A-Z]+(-[a-z0-9A-Z]+)?\\.)+[a-zA-Z]{2,}$", email)
		if !ok {
			flash.Error("请输入正确的邮箱地址")
			flash.Store(&c.Controller)
			c.Redirect("/user/setting", 302)
			return
		}
	}
	if len(signature) > 1000 {
		flash.Error("个人签名长度不能超过1000字符")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}
	_, user := filters.IsLogin(c.Ctx)
	user.Email = email
	user.Signature = signature
	models.UserManager.UpdateUser(&user)
	flash.Success("更新资料成功")
	flash.Store(&c.Controller)
	c.Redirect("/user/setting", 302)
	return
}

func (c *UserController) UpdatePwd() {
	flash := beego.NewFlash()
	oldpwd, newpwd := c.Input().Get("oldpwd"), c.Input().Get("newpwd")
	_, user := filters.IsLogin(c.Ctx)
	if !models.UserManager.CheckPwd(user.Password, oldpwd) {
		flash.Error("旧密码不正确")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}
	if len(newpwd) == 0 {
		flash.Error("新密码不能为空")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}
	user.Password = user.EncodePwd(newpwd)
	models.UserManager.UpdateUser(&user)
	flash.Success("密码修改成功")
	flash.Store(&c.Controller)
	c.Redirect("/user/setting", 302)
	return
}

func (c *UserController) UpdateAvatar() {
	flash := beego.NewFlash()
	f, h, err := c.GetFile("avatar")
	if err == http.ErrMissingFile {
		flash.Error("请选择图片")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}
	defer f.Close()
	if err != nil {
		flash.Error("上传失败")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	} else {

		_, user := filters.IsLogin(c.Ctx)

		dirFile := fmt.Sprintf("%s/%s/%s/%s", beego.AppConfig.String("dirpath"),
			user.Username, "avatar", h.Filename)

		err := c.SaveToFile("avatar", dirFile)
		if err != nil {
			fmt.Printf("\n upload avatar error[%s] \n", err.Error())
			flash.Error("上传失败")
			flash.Store(&c.Controller)
			c.Redirect("/user/setting", 302)
			return
		}

		user.Image = strings.TrimPrefix(user.Image, "/")

		if !strings.Contains(user.Image, "default") {
			err = os.Remove(user.Image) //删除旧头像
			if err != nil {
				fmt.Printf("\n update avatar and delete old avatar error[%s] \n", err.Error())
			}
		}

		user.Image = "/" + dirFile
		models.UserManager.UpdateUser(&user)
		flash.Success("上传成功")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}
}

func (c *UserController) List() {
	beego.ReadFromRequest(&c.Controller)
	c.Data["PageTitle"] = "用户列表"
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
	page, _ := strconv.Atoi(c.Ctx.Input.Query("page"))
	searchName := c.Input().Get("searchName")

	if page == 0 {
		page = 1
	}
	size, _ := beego.AppConfig.Int("page.size")
	c.Data["Page"] = models.UserManager.PageUser(page, size, searchName)
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/list.tpl"
}

func (c *UserController) Edit() {
	c.Data["PageTitle"] = "配置角色"
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if id > 0 {
		ok, user := models.UserManager.FindUserByID(id)
		if ok {
			c.Data["User"] = user
			c.Data["Roles"] = models.RoleManager.FindRoles()
			c.Data["UserRoles"] = models.RoleManager.FindRolesByUser(&user)
			c.Layout = "layout/layout.tpl"
			c.TplName = "user/edit.tpl"
		} else {
			c.Ctx.WriteString("用户不存在")
		}
	} else {
		c.Ctx.WriteString("用户不存在")
	}
}

func (c *UserController) Update() {
	flash := beego.NewFlash()
	c.Data["PageTitle"] = "配置角色"
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	roleIds := c.GetStrings("roleIds")
	if id > 0 {
		models.UserManager.DeleteUserRolesByUserId(id)
		for _, v := range roleIds {
			roleId, _ := strconv.Atoi(v)
			models.UserManager.SaveUserRole(id, roleId)
		}
		flash.Success("修改成功")
		flash.Store(&c.Controller)
		c.Redirect("/user/list", 302)
		return
	} else {
		c.Ctx.WriteString("用户不存在")
	}
}

func (c *UserController) Delete() {
	flash := beego.NewFlash()
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if id > 0 {
		ok, user := models.UserManager.FindUserByID(id)
		if ok {
			models.UserManager.DeleteUser(&user)
			models.UserManager.DeleteUserRolesByUserId(user.Id)

			deletePath := fmt.Sprintf("%s/%s", beego.AppConfig.String("dirpath"), user.Username)
			err := os.RemoveAll(deletePath)
			if err != nil {
				fmt.Printf("\n delete user and delete user dir error[%s] \n", err.Error())
				flash.Error("删除用户时发生错误")
				flash.Store(&c.Controller)
				c.Redirect("/user/list", 302)
				return
			}
		}
		flash.Success("删除成功")
		flash.Store(&c.Controller)
		c.Redirect("/user/list", 302)
		return
	} else {
		c.Ctx.WriteString("用户不存在")
	}
}
