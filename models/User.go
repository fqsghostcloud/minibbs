package models

import (
	"minibbs/utils"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

// API user api
type API interface {
	Login(username string, password string) (bool, User)

	FindUserByID(id int) (bool, User)
	FindUserByToken(token string) (bool, User)
	FindUserByUserName(username string) (bool, User)
	FindPermissionByUser(id int) []*Permission
	FindUserRolesByUserID(userID int) []orm.Params

	SaveUser(user *User) int64
	SaveUserRole(userID int, roleID int)
	DeleteUser(user *User)
	UpdateUser(user *User)
	DeleteUserRolesByUserID(userID int)

	PageUser(p int, size int) utils.Page
}

// User ..
type User struct {
	ID        int    `orm:"pk;auto"`
	Username  string `orm:"unique"`
	Password  string
	Token     string `orm:"unique"`
	Image     string
	Email     string    `orm:"null"`
	URL       string    `orm:"null"`
	Signature string    `orm:"null;size(1000)"`
	InTime    time.Time `orm:"auto_now_add;type(datetime)"`
	Roles     []*Role   `orm:"rel(m2m)"`
}

// UserManager manager user api
var UserManager API

func init() {
	UserManager = new(User)
}

// FindPermissionByUserIDAndPermissionName .
func FindPermissionByUserIDAndPermissionName(userID int, name string) bool {
	o := orm.NewOrm()
	var permission Permission
	o.Raw("select p.* from permission p "+
		"left join role_permissions rp on p.id = rp.permission_id "+
		"left join role r on rp.role_id = r.id "+
		"left join user_roles ur on r.id = ur.role_id "+
		"left join user u on ur.user_id = u.id "+
		"where u.id = ? and p.name = ?", userID, name).QueryRow(&permission)
	return permission.Id > 0
}

// FindUserByID .
func (u *User) FindUserByID(ID int) (bool, User) {
	o := orm.NewOrm()
	err := o.QueryTable(*u).Filter("Id", ID).One(u)
	return err != orm.ErrNoRows, *u
}

// FindUserByToken .
func (u *User) FindUserByToken(token string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Token", token).One(&user)
	return err != orm.ErrNoRows, user
}

// Login .
func (u *User) Login(username string, password string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Username", username).Filter("Password", password).One(&user)
	return err != orm.ErrNoRows, user
}

// FindUserByUserName .
func (u *User) FindUserByUserName(username string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Username", username).One(&user)
	return err != orm.ErrNoRows, user
}

// SaveUser .
func (u *User) SaveUser(user *User) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(user)
	return id
}

// UpdateUser .
func (u *User) UpdateUser(user *User) {
	o := orm.NewOrm()
	o.Update(user)
}

// PageUser .
func (u *User) PageUser(p int, size int) utils.Page {
	o := orm.NewOrm()
	var user User
	var list []User
	qs := o.QueryTable(user)
	count, _ := qs.Limit(-1).Count()
	qs.RelatedSel().OrderBy("-InTime").Limit(size).Offset((p - 1) * size).All(&list)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.PageUtil(c, p, size, list)
}

// FindPermissionByUser .
func (u *User) FindPermissionByUser(ID int) []*Permission {
	o := orm.NewOrm()
	var permissions []*Permission
	o.Raw("select p.* from permission p "+
		"left join role_permissions rp on p.id = rp.permission_id "+
		"left join role r on rp.role_id = r.id "+
		"left join user_roles ur on r.id = ur.role_id "+
		"left join user u on ur.user_id = u.id "+
		"where u.id = ?", ID).QueryRows(&permissions)
	return permissions
}

// DeleteUser .
func (u *User) DeleteUser(user *User) {
	o := orm.NewOrm()
	o.Delete(user)
}

// DeleteUserRolesByUserID .
func (u *User) DeleteUserRolesByUserID(userID int) {
	o := orm.NewOrm()
	o.Raw("delete from user_roles where user_id = ?", userID).Exec()
}

// SaveUserRole .
func (u *User) SaveUserRole(userID int, roleID int) {
	o := orm.NewOrm()
	o.Raw("insert into user_roles (user_id, role_id) values (?, ?)", userID, roleID).Exec()
}

// FindUserRolesByUserID .
func (u *User) FindUserRolesByUserID(userID int) []orm.Params {
	o := orm.NewOrm()
	var res []orm.Params
	o.Raw("select id, user_id, role_id from user_roles where user_id = ?", userID).Values(&res, "id", "user_id", "role_id")
	return res
}
