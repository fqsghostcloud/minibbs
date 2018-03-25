package models

import "github.com/astaxie/beego/orm"

type Role struct {
	Id          int           `orm:"pk;auto"`
	Name        string        `orm:"unique"`
	Users       []*User       `orm:"reverse(many)"`
	Permissions []*Permission `orm:"rel(m2m)"`
}

type RoleAPI interface {
	FindRoleById(id int) Role
	FindRoles() []*Role
	SaveRole(role *Role) int64
	UpdateRole(role *Role)
	DeleteRole(role *Role)
	DeleteRolePermissionByRoleId(role_id int)
	SaveRolePermission(role_id int, permission_id int)
	FindRolePermissionByRoleId(role_id int) []orm.Params
}

// RoleManager manager role api
var RoleManager RoleAPI

func init() {
	RoleManager = new(Role)
}

// FindRoleById .
func (r *Role) FindRoleById(id int) Role {
	o := orm.NewOrm()
	var role Role
	o.QueryTable(role).Filter("Id", id).One(&role)
	return role
}

// FindRoles .
func (r *Role) FindRoles() []*Role {
	o := orm.NewOrm()
	var role Role
	var roles []*Role
	o.QueryTable(role).All(&roles)
	return roles
}

// SaveRole .
func (r *Role) SaveRole(role *Role) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(role)
	return id
}

// UpdateRole .
func (r *Role) UpdateRole(role *Role) {
	o := orm.NewOrm()
	o.Update(role, "Name")
}

// DeleteRole .
func (r *Role) DeleteRole(role *Role) {
	o := orm.NewOrm()
	o.Delete(role)
}

// DeleteRolePermissionByRoleId .
func (r *Role) DeleteRolePermissionByRoleId(role_id int) {
	o := orm.NewOrm()
	o.Raw("delete from role_permissions where role_id = ?", role_id).Exec()
}

// SaveRolePermission .
func (r *Role) SaveRolePermission(role_id int, permission_id int) {
	o := orm.NewOrm()
	o.Raw("insert into role_permissions (role_id, permission_id) values (?, ?)", role_id, permission_id).Exec()
}

// FindRolePermissionByRoleId .
func (r *Role) FindRolePermissionByRoleId(role_id int) []orm.Params {
	o := orm.NewOrm()
	var res []orm.Params
	o.Raw("select id, role_id, permission_id from role_permissions where role_id = ?", role_id).Values(&res, "id", "role_id", "permission_id")
	return res
}
