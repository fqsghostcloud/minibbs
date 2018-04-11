package controllers

import (
	"minibbs/filters"
	"minibbs/models"
	"minibbs/utils"
	"strconv"

	"github.com/astaxie/beego"
)

type ReplyController struct {
	beego.Controller
}

func (c *ReplyController) Save() {
	content := c.Input().Get("content")
	if len(content) == 0 {
		c.Ctx.WriteString("回复内容不能为空")
	} else {
		tid, _ := strconv.Atoi(c.Input().Get("tid"))
		if tid == 0 {
			c.Ctx.WriteString("回复的话题不存在")
		} else {
			_, user := filters.IsLogin(c.Ctx)
			topic := models.TopicManager.FindTopicById(tid)
			reply := models.Reply{Content: content, Topic: &topic, User: &user, Up: 0}
			models.ReplyManager.SaveReply(&reply)
			models.TopicManager.IncrReplyCount(&topic)
			c.Redirect("/topic/"+strconv.Itoa(tid), 302)
		}
	}
}

func (c *ReplyController) Up() {
	rid, _ := strconv.Atoi(c.Ctx.Input.Query("rid"))
	result := utils.Result{Code: 200, Description: "成功"}
	if rid > 0 {
		_, user := filters.IsLogin(c.Ctx)
		reply := models.ReplyManager.FindReplyById(rid)
		replyUpLog := models.FindReplyUpLogByUserAndReply(&user, &reply)
		if replyUpLog.Id > 0 {
			replyUpLog.User = &user
			replyUpLog.Reply = &reply
			models.DeleteReplyUpLog(&replyUpLog)
			models.ReplyManager.DownReply(&reply)
			result.Code = 201
			// result.Description = "你已取消点赞"
		} else {
			replyUpLog.User = &user
			replyUpLog.Reply = &reply
			models.SaveReplyUpLog(&replyUpLog)
			models.ReplyManager.UpReply(&reply)
		}
	} else {
		result.Code = 201
		result.Description = "失败1"
	}
	c.Data["json"] = &result
	c.ServeJSON()
}

func (c *ReplyController) Delete() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if id > 0 {
		reply := models.ReplyManager.FindReplyById(id)
		tid := reply.Topic.Id
		models.TopicManager.ReduceReplyCount(reply.Topic)
		models.ReplyManager.DeleteReply(&reply)
		c.Redirect("/topic/"+strconv.Itoa(tid), 302)
	} else {
		c.Ctx.WriteString("回复不存在")
	}
}

func (c *ReplyController) UserReplay() {
	username := c.Ctx.Input.Param(":username")
	size, _ := beego.AppConfig.Int("page.size")
	pageNum, _ := strconv.Atoi(c.Ctx.Input.Query("page"))
	if pageNum == 0 {
		pageNum = 1
	}

	ok, user := models.UserManager.FindUserByUserName(username)
	if ok {
		c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
		c.Data["PageTitle"] = "个人主页"
		c.Data["CurrentUserInfo"] = user
		c.Data["Page"] = models.ReplyManager.FindReplyByUser(&user, -1, pageNum, size)
	}
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/allReply.tpl"
}
