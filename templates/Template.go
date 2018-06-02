package utils

import (
	"minibbs/models"
	"minibbs/utils"
	"time"

	"github.com/astaxie/beego"
	"github.com/russross/blackfriday"
	"github.com/xeonx/timeago"
)

func FormatTime(time time.Time) string {
	return timeago.Chinese.Format(time)
}

func Markdown(content string) string {
	return string(blackfriday.MarkdownCommon([]byte(utils.NoHtml(content))))
}

// HasPermission .
func HasPermission(userID int, name string) bool {
	return models.UserManager.FindPermissionByUserIDAndPermissionName(userID, name)
}

func IsTopicUser(userId int, topicId int) bool {
	user := models.User{Id: userId}

	if !IsAdmin(user) {
		topicUser := models.TopicManager.FindTopicById(topicId).User
		if userId == topicUser.Id {
			return true
		}
		return false
	}

	return true
}

func GetTopicTags(topicId int) []models.Tag {
	topic := models.TopicManager.FindTopicById(topicId)
	return models.TagManager.FindTagsByTopic(&topic)
}

func GetTopicUser(topicId int) string {
	isExsit, user := models.UserManager.FindUserByTopicId(topicId)
	if isExsit {
		return user.Username
	}
	return ""
}

func IsAdmin(user models.User) bool {
	roles := models.RoleManager.FindRolesByUser(&user)
	for _, v := range roles {
		if v.Name == models.ADMIN {
			return true
		}
	}
	return false
}

func HasFile(topicFile string) bool {
	return len(topicFile) > 0
}

func init() {
	beego.AddFuncMap("timeago", FormatTime)
	beego.AddFuncMap("markdown", Markdown)
	beego.AddFuncMap("haspermission", HasPermission)
	beego.AddFuncMap("getTopicTags", GetTopicTags)
	beego.AddFuncMap("isAdmin", IsAdmin)
	beego.AddFuncMap("isTopicUser", IsTopicUser)
	beego.AddFuncMap("hasFile", HasFile)
	beego.AddFuncMap("getTopicUser", GetTopicUser)
}
