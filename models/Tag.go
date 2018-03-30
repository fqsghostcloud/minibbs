package models

import (
	"github.com/astaxie/beego/orm"
)

type Tag struct {
	Id     int      `orm:"pk;auto"`
	Name   string   `orm:"unique"`
	Topics []*Topic `orm:"reverse(many)"`
}

func FindAllTag() []*Tag {
	o := orm.NewOrm()
	var tag Tag
	var tags []*Tag
	o.QueryTable(tag).All(&tags)
	return tags
}
