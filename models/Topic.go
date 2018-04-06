package models

import (
	"fmt"
	"minibbs/utils"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

// TopicAPI api for topic
type TopicAPI interface {
	SaveTopic(topic *Topic, tagsId []*Tag) int64
	FindTopicById(id int) Topic
	SetTagsToTopic(topic *Topic) *Topic
	PageTopic(p int, size int, tag *Tag) utils.Page
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
	View          int       `orm:"default(0)"`
	ReplyCount    int       `orm:"default(0)"`
	LastReplyUser *User     `orm:"rel(fk);null"`
	LastReplyTime time.Time `orm:"auto_now_add;type(datetime)"`
	Tags          []*Tag    `orm:"rel(m2m)"`
}

// TopicManager manager topic api
var TopicManager TopicAPI

func init() {
	TopicManager = new(Topic)
}

// SaveTopic .
func (t *Topic) SaveTopic(topic *Topic, tags []*Tag) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(topic)
	if err != nil {
		fmt.Printf("save topic error: %s", err.Error())
		return -1
	}

	//manytomany
	m2m := o.QueryM2M(topic, "Tags")
	_, err = m2m.Add(tags)
	if err != nil {
		fmt.Printf("save topic error: %s", err.Error())
		return -1
	}

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
func (t *Topic) PageTopic(p int, size int, tag *Tag) utils.Page {
	o := orm.NewOrm()
	var topic Topic
	var list []Topic

	qs := o.QueryTable(topic)
	if tag.Id > 0 {
		qs = qs.Filter("Tags__Tag__Id", tag.Id)
	}
	countStr, _ := qs.Limit(-1).Count()
	qs.RelatedSel().OrderBy("-InTime").Limit(size).Offset((p - 1) * size).All(&list)

	// project pointer problem----bug-------------------------------------
	for k, topic := range list {
		var tags []Tag
		_, err := o.QueryTable(tag).Filter("Topics__Topic__Id", topic.Id).All(&tags)
		if err != nil {
			fmt.Printf("get page topic error[%s]", err.Error())
		}

		for _, ptag := range tags {
			topic.Tags = append(topic.Tags, &ptag)
		}

		list[k] = topic //!!!
	}

	count, _ := strconv.Atoi(strconv.FormatInt(countStr, 10))
	return utils.PageUtil(count, p, size, list)
}

func (t *Topic) SetTagsToTopic(topic *Topic) *Topic {
	o := orm.NewOrm()
	var tags []Tag

	_, err := o.QueryTable(Tag{}).Filter("Topics__Topic__Id", topic.Id).All(&tags)
	if err != nil {
		fmt.Printf("get page topic error[%s]", err.Error())
		return nil
	}

	for _, ptag := range tags {
		topic.Tags = append(topic.Tags, &ptag)
	}

	return topic

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

	_, err := o.QueryTable(topic).RelatedSel().Filter("User", user).OrderBy("-LastReplyTime", "-InTime").Limit(limit).All(&topics)
	if err != nil {
		fmt.Printf("find topic by user error:[%s]", err.Error())
	}

	for _, topic := range topics {
		topic = t.SetTagsToTopic(topic)
	}
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
