package controllers

import (
	"minibbs/filters"
	"minibbs/models"
	"strconv"

	"github.com/astaxie/beego"
)

type RoleController struct {
	beego.Controller
}

func (c *RoleController) List() {
	c.Data["PageTitle"] = "角色列表"
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
	c.Data["Roles"] = models.RoleManager.FindRoles()
	c.Layout = "layout/layout.tpl"
	c.TplName = "role/list.tpl"
}

func (c *RoleController) Add() {
	beego.ReadFromRequest(&c.Controller)
	c.Data["PageTitle"] = "添加角色"
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
	permissions := models.PermissionManager.FindPermissionsByPid(0)
	for _, p := range permissions {
		p.ChildPermissions = models.PermissionManager.FindPermissionsByPid(p.Id)
	}
	c.Data["Permissions"] = permissions
	c.Layout = "layout/layout.tpl"
	c.TplName = "role/add.tpl"
}

func (c *RoleController) Save() {
	flash := beego.NewFlash()
	name, permissionIds := c.GetString("name"), c.GetStrings("permissionIds")
	if len(name) == 0 {
		flash.Error("角色名称不能为空")
		flash.Store(&c.Controller)
		c.Redirect("/role/add", 302)
	} else {
		role := models.Role{Name: name}
		id := models.RoleManager.SaveRole(&role)
		role_id, _ := strconv.Atoi(strconv.FormatInt(id, 10))
		for _, pid := range permissionIds {
			_pid, _ := strconv.Atoi(pid)
			models.RoleManager.SaveRolePermission(role_id, _pid)
		}
		c.Redirect("/role/list", 302)
	}
}

func (c *RoleController) Edit() {
	beego.ReadFromRequest(&c.Controller)
	c.Data["PageTitle"] = "编辑角色"
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if id > 0 {
		c.Data["Role"] = models.RoleManager.FindRoleById(id)
		c.Data["RolePermissions"] = models.RoleManager.FindRolePermissionByRoleId(id)
		permissions := models.PermissionManager.FindPermissionsByPid(0)
		for _, p := range permissions {
			p.ChildPermissions = models.PermissionManager.FindPermissionsByPid(p.Id)
		}
		c.Data["Permissions"] = permissions
		c.Layout = "layout/layout.tpl"
		c.TplName = "role/edit.tpl"
	} else {
		c.Ctx.WriteString("角色不存在")
	}
}

func (c *RoleController) Update() {
	flash := beego.NewFlash()
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	name, permissionIds := c.GetString("name"), c.GetStrings("permissionIds")
	if len(name) == 0 {
		flash.Error("角色名称不能为空")
		flash.Store(&c.Controller)
		c.Redirect("/role/add", 302)
	} else {
		role := models.Role{Id: id, Name: name}
		models.RoleManager.UpdateRole(&role)
		models.RoleManager.DeleteRolePermissionByRoleId(id)
		for _, pid := range permissionIds {
			_pid, _ := strconv.Atoi(pid)
			models.RoleManager.SaveRolePermission(id, _pid)
		}
		c.Redirect("/role/list", 302)
	}
}

func (c *RoleController) Delete() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if id > 0 {
		role := models.Role{Id: id}
		models.RoleManager.DeleteRole(&role)
		c.Redirect("/role/list", 302)
	} else {
		c.Ctx.WriteString("角色不存在")
	}
}
