package controllers

import (
	"minibbs/filters"
	"minibbs/models"
	"strconv"

	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Create() {
	beego.ReadFromRequest(&c.Controller)
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
	c.Data["PageTitle"] = "发布话题"
	c.Data["Tags"] = models.TagManager.FindAllTag()
	c.Layout = "layout/layout.tpl"
	c.TplName = "topic/create.tpl"
}

func (c *TopicController) Save() {
	flash := beego.NewFlash()
	title, content := c.Input().Get("title"), c.Input().Get("content")
	tids := c.GetStrings("tids")
	if len(title) == 0 || len(title) > 120 {
		flash.Error("话题标题不能为空且不能超过120个字符")
		flash.Store(&c.Controller)
		c.Redirect("/topic/create", 302)
	} else if len(tids) == 0 {
		flash.Error("请选择话题标签")
		flash.Store(&c.Controller)
		c.Redirect("/topic/create", 302)
	} else {
		var tags []*models.Tag
		for _, strid := range tids {
			id, _ := strconv.Atoi(strid)
			tags = append(tags, &models.Tag{Id: id})
		}

		_, user := filters.IsLogin(c.Ctx)
		topic := models.Topic{Title: title, Content: content, User: &user}
		id := models.TopicManager.SaveTopic(&topic, tags)
		c.Redirect("/topic/"+strconv.FormatInt(id, 10), 302)
	}
}

func (c *TopicController) Detail() {
	id := c.Ctx.Input.Param(":id")
	tid, _ := strconv.Atoi(id)
	if tid > 0 {
		c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
		topic := models.TopicManager.FindTopicById(tid)
		models.TopicManager.IncrView(&topic) //查看+1
		topicTags := models.TagManager.FindTagsByTopic(&topic)
		c.Data["PageTitle"] = topic.Title
		c.Data["Topic"] = topic
		c.Data["TopicTags"] = topicTags
		c.Data["Replies"] = models.ReplyManager.FindReplyByTopic(&topic)
		c.Layout = "layout/layout.tpl"
		c.TplName = "topic/detail.tpl"
	} else {
		c.Ctx.WriteString("话题不存在")
	}
}

func (c *TopicController) Edit() {
	beego.ReadFromRequest(&c.Controller)
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if id > 0 {
		topic := models.TopicManager.FindTopicById(id)
		c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
		c.Data["PageTitle"] = "编辑话题"
		c.Data["Tags"] = models.TagManager.FindAllTag()
		topicTags := models.TagManager.FindTagsByTopic(&topic)
		c.Data["Topic"] = topic
		c.Data["TopicTags"] = topicTags
		c.Layout = "layout/layout.tpl"
		c.TplName = "topic/edit.tpl"
	} else {
		c.Ctx.WriteString("话题不存在")
	}
}

func (c *TopicController) Update() {
	flash := beego.NewFlash()
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	title, content, tid := c.Input().Get("title"), c.Input().Get("content"), c.Input().Get("tid")
	if len(title) == 0 || len(title) > 120 {
		flash.Error("话题标题不能为空且不能超过120个字符")
		flash.Store(&c.Controller)
		c.Redirect("/topic/edit/"+strconv.Itoa(id), 302)
	} else if len(tid) == 0 {
		flash.Error("请选择话题标签")
		flash.Store(&c.Controller)
		c.Redirect("/topic/edit/"+strconv.Itoa(id), 302)
	} else {
		s, _ := strconv.Atoi(tid)
		tag := models.Tag{Id: s}
		topic := models.TopicManager.FindTopicById(id)
		topic.Title = title
		topic.Content = content
		topic.Tags = append(topic.Tags, &tag)
		models.TopicManager.UpdateTopic(&topic)
		c.Redirect("/topic/"+strconv.Itoa(id), 302)
	}
}

func (c *TopicController) Delete() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if id > 0 {
		topic := models.TopicManager.FindTopicById(id)
		models.TopicManager.DeleteTopic(&topic)
		models.ReplyManager.DeleteReplyByTopic(&topic)
		c.Redirect("/", 302)
	} else {
		c.Ctx.WriteString("话题不存在")
	}
}

func (c *TopicController) Manage() {
	c.Data["PageTitle"] = "帖子列表"
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)

	size, _ := beego.AppConfig.Int("page.size")
	pageNum, _ := strconv.Atoi(c.Ctx.Input.Query("pageNum"))
	if pageNum == 0 {
		pageNum = 1
	}
	c.Data["Page"] = models.TopicManager.PageTopicList(pageNum, size)
	c.Layout = "layout/layout.tpl"
	c.TplName = "topic/manage.tpl"
}
