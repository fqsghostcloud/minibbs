package controllers

import (
	"fmt"
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
		isLogin, currUser := filters.IsLogin(c.Controller.Ctx)

		c.Data["IsLogin"] = isLogin
		c.Data["UserInfo"] = currUser
		// isAdmin := false

		// currRoles := models.RoleManager.FindRolesByUser(&currUser)

		// for _, v := range currRoles {
		// 	if v.Name == models.ADMIN || v.Name == models.SUPERADMIN {
		// 		isAdmin = true
		// 	}
		// }

		// if !isAdmin {
		// 	//whether is current user's topic
		// 	topicUser := models.TopicManager.FindTopicById(id).User
		// 	if topicUser.Username != currUser.Username {
		// 		c.Ctx.WriteString("你无权限访问此页面")
		// 	}
		// }

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
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	tids := c.GetStrings("tids")
	if len(title) == 0 || len(title) > 120 {
		flash.Error("话题标题不能为空且不能超过120个字符")
		flash.Store(&c.Controller)
		c.Redirect("/topic/edit/"+strconv.Itoa(id), 302)
	} else if len(tids) == 0 {
		flash.Error("请选择话题标签")
		flash.Store(&c.Controller)
		c.Redirect("/topic/edit/"+strconv.Itoa(id), 302)
	} else {
		models.TopicManager.DeleteTopicTagsByTopicId(id)
		for _, v := range tids {
			tagId, _ := strconv.Atoi(v)
			models.TopicManager.SaveTopicTag(id, tagId)
		}

		topic := models.TopicManager.FindTopicById(id)
		topic.Title = title
		topic.Content = content
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
		_, user := filters.IsLogin(c.Ctx)
		roles := models.RoleManager.FindRolesByUser(&user)

		for _, v := range roles {
			if v.Name == "管理员" {
				c.Redirect("/topic/manage", 302)
				return
			}
		}
		c.Redirect("/", 302)
		return
	} else {
		c.Ctx.WriteString("话题不存在")
	}
	return
}

func (c *TopicController) Manage() {
	c.Data["PageTitle"] = "帖子列表"
	isLogin, userInfo := filters.IsLogin(c.Ctx)
	c.Data["IsLogin"] = isLogin
	c.Data["UserInfo"] = userInfo
	roles := models.RoleManager.FindRolesByUser(&userInfo)
	isAdmin := false

	for _, v := range roles {
		if v.Name == models.ADMIN {
			isAdmin = true
			break
		}
	}

	size, _ := beego.AppConfig.Int("page.size")
	pageNum, _ := strconv.Atoi(c.Ctx.Input.Query("pageNum"))
	if pageNum == 0 {
		pageNum = 1
	}

	if isAdmin {
		c.Data["Page"] = models.TopicManager.PageTopicList(pageNum, size, nil)
	} else {
		c.Data["Page"] = models.TopicManager.PageTopicList(pageNum, size, &userInfo)
	}

	c.Layout = "layout/layout.tpl"
	c.TplName = "topic/manage.tpl"
}

func (c *TopicController) TagManage() {
	c.Data["PageTitle"] = "标签列表"
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)

	size, _ := beego.AppConfig.Int("page.size")
	pageNum, _ := strconv.Atoi(c.Ctx.Input.Query("pageNum"))
	if pageNum == 0 {
		pageNum = 1
	}
	c.Data["Page"] = models.TagManager.PageTagList(pageNum, size)
	c.Layout = "layout/layout.tpl"
	c.TplName = "topic/manageTag.tpl"
}

func (c *TopicController) SaveTag() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()
	tagName := c.Input().Get("tagName")
	if tagName == "" {
		flash.Error("标签不可以为空")
		flash.Store(&c.Controller)
		c.Redirect("/tag/manage/", 302)
		return
	}

	if len(tagName) == 1 {
		flash.Error("标签至少两个字符")
		flash.Store(&c.Controller)
		c.Redirect("/tag/manage/", 302)
		return
	}

	tag := models.Tag{Name: tagName}

	err := models.TagManager.SaveTag(&tag)
	if err != nil {
		fmt.Printf("\n save tag error[%s] \n", err.Error())
		flash.Error("保存标签时发生错误")
		flash.Store(&c.Controller)
	}
	c.Redirect("/tag/manage/", 302)
	return
}

func (c *TopicController) UpdateTag() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()
	tagName := c.Input().Get("tagName")
	id, _ := strconv.Atoi(c.Input().Get("id"))

	if tagName == "" {
		flash.Error("标签不可以为空")
		flash.Store(&c.Controller)
		c.Redirect("/tag/manage/", 302)
		return
	}

	if len(tagName) == 1 {
		flash.Error("标签至少两个字符")
		flash.Store(&c.Controller)
		c.Redirect("/tag/manage/", 302)
		return
	}

	tag, err := models.TagManager.FinTagById(id)
	if err != nil {
		fmt.Printf("\n update tag error[%s] \n", err.Error())
		flash.Error("查询标签时发生错误")
		flash.Store(&c.Controller)
		c.Redirect("/tag/manage/", 302)
		return
	}

	tag.Name = tagName

	err = models.TagManager.UpdateTag(tag)
	if err != nil {
		fmt.Printf("\n update tag error[%s] \n", err.Error())
		flash.Error("修改标签时发生错误")
		flash.Store(&c.Controller)
	}
	c.Redirect("/tag/manage/", 302)
	return
}

func (c *TopicController) DeleteTag() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if id > 0 {
		tag, _ := models.TagManager.FinTagById(id)
		err := models.TagManager.DeleteTag(tag)
		if err != nil {
			fmt.Printf("\n delete tag error[%s] \n", err.Error())
			flash.Error("删除标签时发生错误")
			flash.Store(&c.Controller)
			c.Redirect("/tag/manage/", 302)
			return
		}

		flash.Success("删除成功")
		flash.Store(&c.Controller)
		c.Redirect("/tag/manage/", 302)
		return

	} else {
		c.Ctx.WriteString("标签不存在")
	}
	return
}

func (c *TopicController) UserTopic() {
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
		c.Data["Page"] = models.TopicManager.FindTopicByUser(&user, -1, pageNum, size)
	}
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/allTopic.tpl"
}

func (c *TopicController) TopicApproval() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()
	topicId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	if topicId < 0 {
		flash.Error("帖子不存在")
		flash.Store(&c.Controller)
		c.Redirect("/topic/manage/", 302)
		return
	}

	topic := models.TopicManager.FindTopicById(topicId)
	if topic.IsApproval == true {
		flash.Notice("帖子已经审核通过")
		flash.Store(&c.Controller)
		c.Redirect("/topic/manage/", 302)
		return
	}

	topic.IsApproval = true

	models.TopicManager.UpdateTopic(&topic)
	flash.Success("审核通过")
	flash.Store(&c.Controller)
	c.Redirect("/topic/manage/", 302)
	return

}

func (c *TopicController) TopicNotApproval() {
	beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()
	topicId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	if topicId < 0 {
		flash.Error("帖子不存在")
		flash.Store(&c.Controller)
		c.Redirect("/topic/manage/", 302)
		return
	}

	topic := models.TopicManager.FindTopicById(topicId)
	if topic.IsApproval == false {
		flash.Notice("审核已经打回")
		flash.Store(&c.Controller)
		c.Redirect("/topic/manage/", 302)
		return
	}

	topic.IsApproval = false

	models.TopicManager.UpdateTopic(&topic)
	flash.Success("审核打回成功")
	flash.Store(&c.Controller)
	c.Redirect("/topic/manage/", 302)
	return

}
