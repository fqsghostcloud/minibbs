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

func GetTopicTags(topicName string) []models.Tag {
	topic := models.TopicManager.FindTopicByName(topicName)
	return models.TagManager.FindTagsByTopic(&topic)
}

func init() {
	beego.AddFuncMap("timeago", FormatTime)
	beego.AddFuncMap("markdown", Markdown)
	beego.AddFuncMap("haspermission", HasPermission)
	beego.AddFuncMap("getTopicTags", GetTopicTags)
}
