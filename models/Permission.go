package models

import "github.com/astaxie/beego/orm"

type PermissionAPI interface {
	FindPermissionById(id int) Permission
	FindPermissions() []*Permission
	FindPermissionsByPid(pid int) []*Permission
	SavePermission(permission *Permission) int64
	UpdatePermission(permission *Permission) int64
	DeletePermission(permission *Permission)
	DeleteRolePermissionByPermissionId(permission_id int)
}

type Permission struct {
	Id               int `orm:"pk;auto"`
	Pid              int
	Url              string
	Name             string
	Description      string
	Roles            []*Role       `orm:"reverse(many)"`
	ChildPermissions []*Permission `orm:"-"`
}

var PermissionManager PermissionAPI

func init() {
	PermissionManager = new(Permission)
}

func (p *Permission) FindPermissionById(id int) Permission {
	o := orm.NewOrm()
	var permission Permission
	o.QueryTable(permission).Filter("Id", id).One(&permission)
	return permission
}

func (p *Permission) FindPermissions() []*Permission {
	o := orm.NewOrm()
	var permission Permission
	var permissions []*Permission
	o.QueryTable(permission).All(&permissions)
	return permissions
}

func (p *Permission) FindPermissionsByPid(pid int) []*Permission {
	o := orm.NewOrm()
	var permission Permission
	var permissions []*Permission
	o.QueryTable(permission).Filter("Pid", pid).All(&permissions)
	return permissions
}

func (p *Permission) SavePermission(permission *Permission) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(permission)
	return id
}

func (p *Permission) UpdatePermission(permission *Permission) int64 {
	o := orm.NewOrm()
	id, _ := o.Update(permission)
	return id
}

func (p *Permission) DeletePermission(permission *Permission) {
	o := orm.NewOrm()
	o.Delete(permission)
}

func (p *Permission) DeleteRolePermissionByPermissionId(permission_id int) {
	o := orm.NewOrm()
	o.Raw("delete from role_permissions where permission_id = ?", permission_id).Exec()
}
