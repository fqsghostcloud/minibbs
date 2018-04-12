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
	SaveTopicTag(topicId int, tagId int)
	FindTopicById(id int) Topic
	FindTopicByName(name string) Topic
	PageTopic(p int, size int, tag *Tag) utils.Page           // get  topic with tag
	PageTopicList(page int, size int, owner *User) utils.Page // just get topic list without tag
	IncrView(topic *Topic)
	IncrReplyCount(topic *Topic)
	ReduceReplyCount(topic *Topic)
	FindTopicByUser(user *User, limit int, page int, size int) utils.Page
	UpdateTopic(topic *Topic)
	DeleteTopic(topic *Topic)
	DeleteTopicByUser(user *User)
	DeleteTopicTagsByTopicId(topicId int)
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
	IsApproval    bool      `orm:"default(false)"`
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

func (t *Topic) FindTopicByName(name string) Topic {
	o := orm.NewOrm()
	var topic Topic
	o.QueryTable(topic).RelatedSel().Filter("Title", name).One(&topic)
	return topic
}

func (t *Topic) PageTopicList(page int, size int, owner *User) utils.Page {
	o := orm.NewOrm()
	var topic Topic
	var list []Topic

	qs := o.QueryTable(topic)
	var countStr int64
	if owner == nil {
		countStr, _ = qs.Limit(-1).Count()
		qs.RelatedSel().OrderBy("-InTime").Limit(size).Offset((page - 1) * size).All(&list)
	} else {
		countStr, _ = qs.Filter("User", owner).Limit(-1).Count()
		qs.Filter("User", owner).RelatedSel().OrderBy("-InTime").Limit(size).Offset((page - 1) * size).All(&list)
	}

	count, _ := strconv.Atoi(strconv.FormatInt(countStr, 10))
	return utils.PageUtil(count, page, size, list)
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
	countStr, _ := qs.Filter("isApproval", true).Limit(-1).Count()
	qs.Filter("isApproval", true).RelatedSel().OrderBy("-InTime").Limit(size).Offset((p - 1) * size).All(&list)

	count, _ := strconv.Atoi(strconv.FormatInt(countStr, 10))
	return utils.PageUtil(count, p, size, list)
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
func (t *Topic) FindTopicByUser(user *User, limit int, page int, size int) utils.Page {
	o := orm.NewOrm()
	var topic Topic
	var topics []Topic

	//不分页
	if page == 0 || size == 0 {
		_, err := o.QueryTable(topic).RelatedSel().Filter("User", user).OrderBy("-LastReplyTime", "-InTime").Limit(limit).All(&topics)
		if err != nil {
			fmt.Printf("find topic by user error:[%s]", err.Error())
		}

		page := utils.Page{List: topics}

		return page
	}

	qs := o.QueryTable(topic)
	countStr, _ := qs.Limit(-1).Count()
	qs.RelatedSel().OrderBy("-InTime").Limit(size).Offset((page - 1) * size).All(&topics)

	count, _ := strconv.Atoi(strconv.FormatInt(countStr, 10))
	return utils.PageUtil(count, page, size, topics)

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

// SaveUserRole .
func (t *Topic) SaveTopicTag(topicId int, tagId int) {
	o := orm.NewOrm()
	o.Raw("insert into topic_tags (topic_id, tag_id) values (?, ?)", topicId, tagId).Exec()
}

func (t *Topic) DeleteTopicTagsByTopicId(topicId int) {
	o := orm.NewOrm()
	o.Raw("delete from topic_tags where topic_id = ?", topicId).Exec()
}
