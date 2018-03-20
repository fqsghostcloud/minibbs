package models

import (
	"fmt"
	"minibbs/utils"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

// API user api
type API interface {
	Login(username string, password string) (bool, User)
	IsActive(username string) (bool, error)

	ActiveUserByEmail(email string) error
	ActiveAccount(email string) error
	ActiveUser(username string) error
	DeactiveUser(username string) error

	ExsitUser(username string) bool
	FindUserByID(id int) (bool, User)
	FindUserByToken(token string) (bool, User)
	FindUserByUserName(username string) (bool, User)
	FindUserByUserEmail(email string) (bool, User)
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
	Active    bool      `orm:"default(false)"`
	Status    bool      `orm:"default(false)"`
}

// UserManager manager user api
var UserManager API

func init() {
	UserManager = new(User)
}

// ActiveAccount .
func (u *User) ActiveAccount(email string) error {
	return u.ActiveUserByEmail(email)
}

// ExsitUser check user whether exsit
func (u *User) ExsitUser(username string) bool {
	o := orm.NewOrm()
	var user User
	qs := o.QueryTable(user)
	return qs.Filter("Username", username).Exist()
}

// IsActive check user whether activve
func (u *User) IsActive(username string) (bool, error) {
	var user User
	o := orm.NewOrm()
	qs := o.QueryTable(user)
	err := qs.Filter("Username", username).One(&user)

	if err != nil {
		return false, err
	}

	return user.Active == true, nil
}

// GetUserByEmail ..
func (u *User) GetUserByEmail(email string) (string, error) {
	o := orm.NewOrm()

	var user User
	qs := o.QueryTable(user)
	err := qs.Filter("Email", email).One(&user)
	if err != nil {
		return "", err
	}

	return user.Username, nil

}

// ActiveUserByEmail ..
func (u *User) ActiveUserByEmail(email string) error {
	exsit, user := u.FindUserByUserName(email)
	if !exsit {
		return fmt.Errorf("此email[%s]的用户不存在", email)
	}

	err := u.ActiveUser(user.Username)
	if err != nil {
		return err
	}

	return nil
}

// ActiveUser active user
func (u *User) ActiveUser(username string) error {
	if !u.ExsitUser(username) {
		return fmt.Errorf("%s", "用户不存在")
	}

	o := orm.NewOrm()

	isActive, err := u.IsActive(username)
	if err != nil {
		return err
	}

	if isActive {
		err := fmt.Errorf("user[%s] already active", username)
		// glog.Infof(err.Error())

		return err
	}

	var user User

	qs := o.QueryTable(user)
	_, err = qs.Filter("Name", username).Update(orm.Params{"Active": true})
	if err != nil {
		return err
	}
	// glog.Infof("active user[%s] success\n", username)

	return nil
}

// InactiveUserByEmail ..
func (u *User) InactiveUserByEmail(email string) error {
	username, err := u.GetUserByEmail(email)
	if err != nil {
		return err
	}

	err = u.DeactiveUser(username)
	if err != nil {
		return err
	}

	return nil
}

// DeactiveUser inactive user
func (u *User) DeactiveUser(username string) error {
	if !u.ExsitUser(username) {
		return fmt.Errorf("%s", "用户不存在")
	}

	isActive, err := u.IsActive(username)
	if err != nil {
		return err
	}

	if !isActive {
		err := fmt.Errorf("user[%s] already inactive", username)
		return err
	}

	var user User
	o := orm.NewOrm()
	qs := o.QueryTable(user)

	_, err = qs.Filter("Name", username).Update(orm.Params{"Active": false})
	if err != nil {
		return err
	}

	return nil
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

// FindUserByUserEmail .
func (u *User) FindUserByUserEmail(email string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Email", email).One(&user)
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
