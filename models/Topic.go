package models

import (
	"minibbs/utils"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

// TopicAPI api for topic
type TopicAPI interface {
	SaveTopic(topic *Topic) int64
	FindTopicById(id int) Topic
	PageTopic(p int, size int, section *Section) utils.Page
	IncrView(topic *Topic)
	IncrReplyCount(topic *Topic)
	ReduceReplyCount(topic *Topic)
	FindTopicByUser(user *User, limit int) []*Topic
	UpdateTopic(topic *Topic)
	DeleteTopic(topic *Topic)
	DeleteTopicByUser(user *User)
}

type Topic struct {
	Id            int       `orm:"pk;auto"`
	Title         string    `orm:"unique"`
	Content       string    `orm:"type(text);null"`
	InTime        time.Time `orm:"auto_now_add;type(datetime)"`
	User          *User     `orm:"rel(fk)"`
	Section       *Section  `orm:"rel(fk)"`
	View          int       `orm:"default(0)"`
	ReplyCount    int       `orm:"default(0)"`
	LastReplyUser *User     `orm:"rel(fk);null"`
	LastReplyTime time.Time `orm:"auto_now_add;type(datetime)"`
}

// TopicManager manager topic api
var TopicManager TopicAPI

func init() {
	TopicManager = new(Topic)
}

// SaveTopic .
func (t *Topic) SaveTopic(topic *Topic) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(topic)
	return id
}

// FindTopicById .
func (t *Topic) FindTopicById(id int) Topic {
	o := orm.NewOrm()
	var topic Topic
	o.QueryTable(topic).RelatedSel().Filter("Id", id).One(&topic)
	return topic
}

// PageTopic .
func (t *Topic) PageTopic(p int, size int, section *Section) utils.Page {
	o := orm.NewOrm()
	var topic Topic
	var list []Topic
	qs := o.QueryTable(topic)
	if section.Id > 0 {
		qs = qs.Filter("Section", section)
	}
	count, _ := qs.Limit(-1).Count()
	qs.RelatedSel().OrderBy("-InTime").Limit(size).Offset((p - 1) * size).All(&list)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.PageUtil(c, p, size, list)
}

func (t *Topic) IncrView(topic *Topic) {
	o := orm.NewOrm()
	topic.View = topic.View + 1
	o.Update(topic, "View")
}

func (t *Topic) IncrReplyCount(topic *Topic) {
	o := orm.NewOrm()
	topic.ReplyCount = topic.ReplyCount + 1
	o.Update(topic, "ReplyCount")
}

func (t *Topic) ReduceReplyCount(topic *Topic) {
	o := orm.NewOrm()
	topic.ReplyCount = topic.ReplyCount - 1
	o.Update(topic, "ReplyCount")
}

// FindTopicByUser .
func (t *Topic) FindTopicByUser(user *User, limit int) []*Topic {
	o := orm.NewOrm()
	var topic Topic
	var topics []*Topic
	o.QueryTable(topic).RelatedSel().Filter("User", user).OrderBy("-LastReplyTime", "-InTime").Limit(limit).All(&topics)
	return topics
}

// UpdateTopic .
func (t *Topic) UpdateTopic(topic *Topic) {
	o := orm.NewOrm()
	o.Update(topic)
}

// DeleteTopic .
func (t *Topic) DeleteTopic(topic *Topic) {
	o := orm.NewOrm()
	o.Delete(topic)
}

// DeleteTopicByUser .
func (t *Topic) DeleteTopicByUser(user *User) {
	o := orm.NewOrm()
	o.Raw("delete from topic where user_id = ?", user.Id).Exec()
}
