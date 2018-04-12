package models

import (
	"fmt"
	"minibbs/utils"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type Tag struct {
	Id     int      `orm:"pk;auto"`
	Name   string   `orm:"unique"`
	Topics []*Topic `orm:"reverse(many)"`
}

type TagAPI interface {
	FindAllTag(user *User) []Tag
	FindTagsByTopic(topic *Topic) []Tag
	FinTagById(id int) (*Tag, error)
	PageTagList(p int, size int) utils.Page
	SaveTag(tag *Tag) error
	UpdateTag(tag *Tag) error
	DeleteTag(tag *Tag) error
}

// TagManager manager tag api
var TagManager TagAPI

func init() {
	TagManager = new(Tag)
}

func (t *Tag) FindAllTag(user *User) []Tag {
	o := orm.NewOrm()
	var tags []Tag

	if user != nil {
		isAdmin := false
		roles := RoleManager.FindRolesByUser(user)
		for _, v := range roles {
			if v.Name == ADMIN {
				isAdmin = true
				break
			}
		}

		if isAdmin {
			o.QueryTable(Tag{}).All(&tags)
			return tags
		}
		o.QueryTable(Tag{}).Exclude("Name", "公告").All(&tags)
		return tags
	} else {
		o.QueryTable(Tag{}).All(&tags)
		return tags
	}

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

func (t *Tag) PageTagList(p int, size int) utils.Page {
	o := orm.NewOrm()
	var tag Tag
	var list []Tag

	qs := o.QueryTable(tag)
	countStr, _ := qs.Limit(-1).Count()
	qs.RelatedSel().Limit(size).Offset((p - 1) * size).All(&list)

	count, _ := strconv.Atoi(strconv.FormatInt(countStr, 10))
	return utils.PageUtil(count, p, size, list)
}

func (t *Tag) SaveTag(tag *Tag) error {
	o := orm.NewOrm()
	_, err := o.Insert(tag)
	if err != nil {
		return err
	}

	return nil
}

func (t *Tag) UpdateTag(tag *Tag) error {
	o := orm.NewOrm()
	_, err := o.Update(tag, "Name")
	if err != nil {
		return err
	}
	return nil
}

func (t *Tag) DeleteTag(tag *Tag) error {
	o := orm.NewOrm()
	_, err := o.QueryTable(Tag{}).Filter("Id", tag.Id).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (t *Tag) FinTagById(id int) (*Tag, error) {
	o := orm.NewOrm()
	tag := Tag{}
	err := o.QueryTable(tag).Filter("Id", id).One(&tag)
	if err != nil {
		return nil, err
	}
	return &tag, nil
}
