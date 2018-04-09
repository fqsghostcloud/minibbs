package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Tag struct {
	Id     int      `orm:"pk;auto"`
	Name   string   `orm:"unique"`
	Topics []*Topic `orm:"reverse(many)"`
}

type TagAPI interface {
	FindAllTag() []Tag
	FindTagsByTopic(topic *Topic) []Tag
}

// TagManager manager tag api
var TagManager TagAPI

func init() {
	TagManager = new(Tag)
}

func (t *Tag) FindAllTag() []Tag {
	o := orm.NewOrm()

	var tags []Tag
	o.QueryTable(Tag{}).All(&tags)
	return tags
}

// FindTagsByTopic .
func (t *Tag) FindTagsByTopic(topic *Topic) []Tag {
	o := orm.NewOrm()
	var tags []Tag

	_, err := o.QueryTable(Tag{}).Filter("Topics__Topic__Id", topic.Id).All(&tags)
	if err != nil {
		fmt.Printf("get page topic error[%s]", err.Error())
		return nil
	}

	return tags
}
