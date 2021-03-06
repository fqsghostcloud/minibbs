package models

import (
	"fmt"
	"minibbs/utils"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

const (
	ADMIN = "管理员"
	USER  = "普通用户"
)

// API user api
type API interface {
	Login(username string, password string) (bool, *User, error)
	IsActive(username string) bool

	ActiveUserByEmail(email string) error
	ActiveAccount(email string) error
	ActiveUser(username string) error
	DeactiveUser(username string) error

	ExsitUser(username string) bool
	FindUserByID(id int) (bool, User)
	FindUserByToken(token string) (bool, User)
	FindUserByUserName(username string) (bool, User)
	FindUserByUserEmail(email string) (bool, User)
	FindUserByTopicId(topicId int) (bool, User)
	FindPermissionByUser(id int) []*Permission

	SaveUser(user *User) error
	SaveUserRole(userID int, roleID int)
	DeleteUser(user *User)
	UpdateUser(user *User)
	DeleteUserRolesByUserId(userId int)

	PageUser(p int, size int, searchName string) utils.Page
	FindPermissionByUserIDAndPermissionName(userID int, name string) bool

	//密码储存加密
	EncodePwd(pwd string) string
	CheckPwd(encodePwd string, pwd string) bool
}

// User ..
type User struct {
	Id        int    `orm:"pk;auto"`
	Username  string `orm:"unique"`
	Password  string
	Token     string `orm:"unique"`
	Image     string
	Email     string
	Signature string    `orm:"null;size(1000)"`
	InTime    time.Time `orm:"auto_now_add;type(datetime)"`
	Roles     []*Role   `orm:"rel(m2m)"`
	Topics    []*Topic  `orm:"reverse(many)"`
	Active    bool      `orm:"default(false)"`
	Status    bool      `orm:"default(false)"`
	LastTime  string    `orm:"type(datetime)"` // last time to login
}

// UserManager manager user api
var UserManager API

func init() {
	UserManager = new(User)
}

func (u *User) EncodePwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	return string(hash)
}

func (u *User) CheckPwd(encodePwd string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodePwd), []byte(pwd))
	if err == nil {
		return true
	}
	return false
}

func (u *User) FindUserByTopicId(topicId int) (bool, User) {
	o := orm.NewOrm()
	user := User{}
	err := o.QueryTable(User{}).Filter("Topics__Id", topicId).One(&user)
	if err != nil {
		return false, user
	}

	return true, user
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
func (u *User) IsActive(username string) bool {
	var user User
	o := orm.NewOrm()
	qs := o.QueryTable(user)
	err := qs.Filter("Username", username).One(&user)
	if err != nil {
		fmt.Println(err.Error())
	}
	return user.Active == true
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
	exsit, user := u.FindUserByUserEmail(email)
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

	isActive := u.IsActive(username)

	if isActive {
		err := fmt.Errorf("user[%s] already active", username)
		// glog.Infof(err.Error())

		return err
	}

	o := orm.NewOrm()

	var user User

	qs := o.QueryTable(user)
	_, err := qs.Filter("Username", username).Update(orm.Params{"Active": true})
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

	isActive := u.IsActive(username)

	if !isActive {
		err := fmt.Errorf("user[%s] already inactive", username)
		return err
	}

	var user User
	o := orm.NewOrm()
	qs := o.QueryTable(user)

	_, err := qs.Filter("Username", username).Update(orm.Params{"Active": false})
	if err != nil {
		return err
	}

	return nil
}

// FindPermissionByUserIDAndPermissionName .
func (u *User) FindPermissionByUserIDAndPermissionName(userID int, name string) bool {
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
func (u *User) FindUserByID(Id int) (bool, User) {
	o := orm.NewOrm()
	err := o.QueryTable(*u).Filter("Id", Id).One(u)
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
func (u *User) Login(username string, password string) (bool, *User, error) {

	if !u.ExsitUser(username) {
		return false, nil, fmt.Errorf("该用户不存在")
	}

	if !u.IsActive(username) {
		return false, nil, fmt.Errorf("该用户帐号未激活")
	}

	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Username", username).One(&user)
	if err != nil {
		fmt.Printf("\n query user error[%s] \n", err.Error())
	}

	if UserManager.CheckPwd(user.Password, password) {
		return true, &user, nil
	}

	return false, &user, fmt.Errorf("用户名或密码错误")
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
func (u *User) SaveUser(user *User) error {

	user.Password = u.EncodePwd(user.Password) // 密码储存加密
	o := orm.NewOrm()
	_, err := o.Insert(user)
	return err
}

// UpdateUser .
func (u *User) UpdateUser(user *User) {
	o := orm.NewOrm()
	o.Update(user)
}

// PageUser .
func (u *User) PageUser(p int, size int, searchName string) utils.Page {
	o := orm.NewOrm()
	var user User
	var list []User
	qs := o.QueryTable(user)
	if searchName == "" {
		count, _ := qs.Limit(-1).Count()
		qs.RelatedSel().OrderBy("-InTime").Limit(size).Offset((p - 1) * size).All(&list)
		c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
		return utils.PageUtil(c, p, size, list)
	} else {
		count, _ := qs.Filter("Username__contains", searchName).Limit(-1).Count()
		qs.Filter("Username__contains", searchName).RelatedSel().OrderBy("-InTime").Limit(size).Offset((p - 1) * size).All(&list)
		c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
		return utils.PageUtil(c, p, size, list)
	}

}

// FindPermissionByUser .
func (u *User) FindPermissionByUser(Id int) []*Permission {
	o := orm.NewOrm()
	var permissions []*Permission
	o.Raw("select p.* from permission p "+
		"left join role_permissions rp on p.id = rp.permission_id "+
		"left join role r on rp.role_id = r.id "+
		"left join user_roles ur on r.id = ur.role_id "+
		"left join user u on ur.user_id = u.id "+
		"where u.id = ?", Id).QueryRows(&permissions)
	return permissions
}

// DeleteUser .
func (u *User) DeleteUser(user *User) {
	o := orm.NewOrm()
	_, err := o.QueryTable(User{}).Filter("Id", user.Id).Delete()
	if err != nil {
		fmt.Printf("\n delete user error[%s] \n", err.Error())
	}
}

// DeleteUserRolesByUserID .
func (u *User) DeleteUserRolesByUserId(userId int) {
	o := orm.NewOrm()
	o.Raw("delete from user_roles where user_id = ?", userId).Exec()
}

// SaveUserRole .
func (u *User) SaveUserRole(userID int, roleID int) {
	o := orm.NewOrm()
	o.Raw("insert into user_roles (user_id, role_id) values (?, ?)", userID, roleID).Exec()
}

// FindUserRolesByUserID .
// func (u *User) FindRolesByUser(user *User) []Role {
// 	o := orm.NewOrm()
// 	var res []orm.Params
// 	o.Raw("select id, user_id, role_id from user_roles where user_id = ?", userID).Values(&res, "id", "user_id", "role_id")
// 	return res
// }
