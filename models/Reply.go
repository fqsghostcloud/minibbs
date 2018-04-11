package models

import (
	"minibbs/utils"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type ReplyAPI interface {
	FindReplyById(id int) Reply
	FindReplyByTopic(topic *Topic) []*Reply
	SaveReply(reply *Reply) int64
	UpReply(reply *Reply)
	FindReplyByUser(user *User, limit int, page int, size int) utils.Page
	DeleteReplyByTopic(topic *Topic)
	DeleteReply(reply *Reply)
	DeleteReplyByUser(user *User)
}

type Reply struct {
	Id      int       `orm:"pk;auto"`
	Topic   *Topic    `orm:"rel(fk)"`
	Content string    `orm:"type(text)"`
	User    *User     `orm:"rel(fk)"`
	Up      int       `orm:"default(0)"`
	InTime  time.Time `orm:"auto_now_add;type(datetime)"`
}

var ReplyManager ReplyAPI

func init() {
	ReplyManager = new(Reply)
}
func (r *Reply) FindReplyById(id int) Reply {
	o := orm.NewOrm()
	var reply Reply
	o.QueryTable(reply).RelatedSel("Topic").Filter("Id", id).One(&reply)
	return reply
}

func (r *Reply) FindReplyByTopic(topic *Topic) []*Reply {
	o := orm.NewOrm()
	var reply Reply
	var replies []*Reply
	o.QueryTable(reply).RelatedSel().Filter("Topic", topic).OrderBy("-Up", "-InTime").All(&replies)
	return replies
}

func (r *Reply) SaveReply(reply *Reply) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(reply)
	return id
}

func (r *Reply) UpReply(reply *Reply) {
	o := orm.NewOrm()
	reply.Up = reply.Up + 1
	o.Update(reply, "Up")
}

func (r *Reply) FindReplyByUser(user *User, limit int, page int, size int) utils.Page {
	o := orm.NewOrm()
	var reply Reply
	var replies []Reply

	if page == 0 || size == 0 {
		o.QueryTable(reply).RelatedSel("Topic", "User").Filter("User", user).OrderBy("-InTime").Limit(limit).All(&replies)
		page := utils.Page{List: replies}
		return page
	}

	qs := o.QueryTable(reply)
	countStr, _ := qs.Limit(-1).Count()
	qs.RelatedSel().OrderBy("-InTime").Limit(size).Offset((page - 1) * size).All(&replies)

	count, _ := strconv.Atoi(strconv.FormatInt(countStr, 10))
	return utils.PageUtil(count, page, size, replies)

}

func (r *Reply) DeleteReplyByTopic(topic *Topic) {
	o := orm.NewOrm()
	var reply Reply
	var replies []Reply
	o.QueryTable(reply).Filter("Topic", topic).All(&replies)
	for _, reply := range replies {
		o.Delete(&reply)
	}
}

func (r *Reply) DeleteReply(reply *Reply) {
	o := orm.NewOrm()
	o.Delete(reply)
}

func (r *Reply) DeleteReplyByUser(user *User) {
	o := orm.NewOrm()
	o.Raw("delete form reply where user_id = ?", user.Id).Exec()
}
