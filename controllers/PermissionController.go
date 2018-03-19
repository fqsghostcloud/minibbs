package controllers

import (
	"minibbs/filters"
	"minibbs/models"
	"strconv"

	"github.com/astaxie/beego"
)

type PermissionController struct {
	beego.Controller
}

func (c *PermissionController) List() {
	c.Data["PageTitle"] = "权限列表"
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
	pid, _ := strconv.Atoi(c.Ctx.Input.Query("pid"))
	if pid > 0 {
		c.Data["Permissions"] = models.PermissionManager.FindPermissionsByPid(pid)
		c.Data["Pid"] = pid
	} else {
		c.Data["Permissions"] = models.PermissionManager.FindPermissions()
	}
	c.Data["ParantPermissions"] = models.PermissionManager.FindPermissionsByPid(0)
	c.Layout = "layout/layout.tpl"
	c.TplName = "permission/list.tpl"
}

func (c *PermissionController) Add() {
	beego.ReadFromRequest(&c.Controller)
	c.Data["PageTitle"] = "添加权限"
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
	c.Data["Pid"] = c.Ctx.Input.Query("pid")
	c.Data["ParantPermissions"] = models.PermissionManager.FindPermissionsByPid(0)
	c.Layout = "layout/layout.tpl"
	c.TplName = "permission/add.tpl"
}

func (c *PermissionController) Save() {
	flash := beego.NewFlash()
	pid, _ := strconv.Atoi(c.Input().Get("pid"))
	name, url, description := c.Input().Get("name"), c.Input().Get("url"), c.Input().Get("description")
	if pid > 0 && len(name) == 0 {
		flash.Error("权限标识不能为空")
		flash.Store(&c.Controller)
		c.Redirect("/permission/add?pid="+strconv.Itoa(pid), 302)
	} else if pid > 0 && len(url) == 0 {
		flash.Error("授权地址不能为空")
		flash.Store(&c.Controller)
		c.Redirect("/permission/add?pid="+strconv.Itoa(pid), 302)
	} else if len(description) == 0 {
		flash.Error("权限描述不能为空")
		flash.Store(&c.Controller)
		c.Redirect("/permission/add?pid="+strconv.Itoa(pid), 302)
	} else {
		permission := models.Permission{Pid: pid, Name: name, Url: url, Description: description}
		models.PermissionManager.SavePermission(&permission)
		c.Redirect("/permission/list?pid="+strconv.Itoa(pid), 302)
	}
}

func (c *PermissionController) Edit() {
	beego.ReadFromRequest(&c.Controller)
	c.Data["PageTitle"] = "编辑权限"
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	c.Data["Permission"] = models.PermissionManager.FindPermissionById(id)
	c.Data["ParantPermissions"] = models.PermissionManager.FindPermissionsByPid(0)
	c.Layout = "layout/layout.tpl"
	c.TplName = "permission/edit.tpl"
}

func (c *PermissionController) Update() {
	flash := beego.NewFlash()
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	pid, _ := strconv.Atoi(c.Input().Get("pid"))
	name, url, description := c.Input().Get("name"), c.Input().Get("url"), c.Input().Get("description")
	if pid > 0 && len(name) == 0 {
		flash.Error("权限标识不能为空")
		flash.Store(&c.Controller)
		c.Redirect("/permission/edit/"+strconv.Itoa(id), 302)
	} else if pid > 0 && len(url) == 0 {
		flash.Error("授权地址不能为空")
		flash.Store(&c.Controller)
		c.Redirect("/permission/edit/"+strconv.Itoa(id), 302)
	} else if len(description) == 0 {
		flash.Error("权限描述不能为空")
		flash.Store(&c.Controller)
		c.Redirect("/permission/edit/"+strconv.Itoa(id), 302)
	} else {
		permission := models.Permission{Id: id, Pid: pid, Name: name, Url: url, Description: description}
		models.PermissionManager.UpdatePermission(&permission)
		c.Redirect("/permission/list?pid="+strconv.Itoa(pid), 302)
	}
}

func (c *PermissionController) Delete() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if id > 0 {
		permission := models.Permission{Id: id}
		models.PermissionManager.DeleteRolePermissionByPermissionId(id)
		models.PermissionManager.DeletePermission(&permission)
		c.Redirect("/permission/list", 302)
	} else {
		c.Ctx.WriteString("权限不存在")
	}
}
