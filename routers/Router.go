package routers

import (
	"minibbs/controllers"
	"minibbs/filters"

	"github.com/astaxie/beego"
)

func init() {

	//登录注册
	beego.Router("/", &controllers.IndexController{}, "GET:Index")
	beego.Router("/login", &controllers.IndexController{}, "GET:LoginPage")
	beego.Router("/login", &controllers.IndexController{}, "POST:Login")
	beego.Router("/register", &controllers.IndexController{}, "GET:RegisterPage")
	beego.Router("/register", &controllers.IndexController{}, "POST:Register")
	beego.Router("/logout", &controllers.IndexController{}, "GET:Logout")

	//聊天室
	beego.InsertFilter("/topic/join/ws/*", beego.BeforeRouter, filters.HasPermission)
	beego.Router("/topic/join/ws", &controllers.ChatRoomController{}, "GET:ChatRoomPage")
	beego.Router("/topic/join/ws/chat", &controllers.ChatRoomController{}, "GET:Chat") // bug

	//创建，修改，删除帖子
	beego.InsertFilter("/topic/create/*", beego.BeforeRouter, filters.HasPermission)
	beego.Router("/topic/create", &controllers.TopicController{}, "GET:Create")
	beego.Router("/topic/create", &controllers.TopicController{}, "POST:Save")

	beego.Router("/topic/:id([0-9]+)", &controllers.TopicController{}, "GET:Detail")
	//下载附件
	beego.InsertFilter("/topic/:id([0-9]+)/download", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/topic/:id([0-9]+)/download", &controllers.TopicController{}, "GET:Download")

	beego.InsertFilter("/topic/edit/:id([0-9]+)", beego.BeforeRouter, filters.HasPermission)
	beego.InsertFilter("/topic/edit/:id([0-9]+)", beego.BeforeRouter, filters.IsTopicUser)
	beego.Router("/topic/edit/:id([0-9]+)", &controllers.TopicController{}, "GET:Edit")
	beego.Router("/topic/edit/:id([0-9]+)", &controllers.TopicController{}, "POST:Update")

	// beego.InsertFilter("/topic/edit/insertpic", beego.BeforeRouter, filters.HasPermission)
	// beego.InsertFilter("/topic/edit/insertpic", beego.BeforeRouter, filters.IsTopicUser)
	beego.Router("/topic/edit/insertpic", &controllers.TopicController{}, "POST:InsertPic")

	beego.InsertFilter("/topic/delete/:id([0-9]+)", beego.BeforeRouter, filters.HasPermission)
	beego.InsertFilter("/topic/delete/:id([0-9]+)", beego.BeforeRouter, filters.IsTopicUser)
	beego.Router("/topic/delete/:id([0-9]+)", &controllers.TopicController{}, "GET:Delete")
	//发帖管理
	beego.InsertFilter("/topic/manage", beego.BeforeRouter, filters.HasPermission)
	beego.Router("/topic/manage", &controllers.TopicController{}, "GET:Manage")
	beego.Router("/topic/manage/:id([0-9]+)/approval", &controllers.TopicController{}, "GET:TopicApproval")
	beego.Router("/topic/manage/:id([0-9]+)/notapproval", &controllers.TopicController{}, "GET:TopicNotApproval")

	//标签管理
	beego.InsertFilter("/tag/manage", beego.BeforeRouter, filters.HasPermission)
	beego.Router("/tag/manage", &controllers.TopicController{}, "GET:TagManage")
	beego.Router("/tag/manage/save", &controllers.TopicController{}, "Post:SaveTag")
	beego.Router("/tag/manage/update", &controllers.TopicController{}, "Post:UpdateTag")
	beego.Router("/tag/manage/delete/:id([0-9]+)", &controllers.TopicController{}, "GET:DeleteTag")

	//回复
	beego.InsertFilter("/reply/save", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/reply/save", &controllers.ReplyController{}, "POST:Save")

	beego.InsertFilter("/reply/up", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/reply/up", &controllers.ReplyController{}, "GET:Up")

	beego.InsertFilter("/reply/delete/:id([0-9]+)", beego.BeforeRouter, filters.HasPermission)
	beego.Router("/reply/delete/:id([0-9]+)", &controllers.ReplyController{}, "GET:Delete")

	//用户
	beego.Router("/user/:username", &controllers.UserController{}, "GET:Detail")
	beego.Router("/user/setting", &controllers.UserController{}, "GET:ToSetting")

	beego.InsertFilter("/user/setting", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/user/setting", &controllers.UserController{}, "POST:Setting")

	beego.InsertFilter("/user/updatepwd", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/user/updatepwd", &controllers.UserController{}, "POST:UpdatePwd")

	beego.InsertFilter("/user/updateavatar", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/user/updateavatar", &controllers.UserController{}, "POST:UpdateAvatar")

	beego.InsertFilter("/user/list", beego.BeforeRouter, filters.HasPermission)
	beego.Router("/user/list", &controllers.UserController{}, "GET:List")

	//用户管理
	beego.InsertFilter("/user/edit/:id([0-9]+)", beego.BeforeRouter, filters.HasPermission)
	beego.Router("/user/edit/:id([0-9]+)", &controllers.UserController{}, "GET:Edit")
	beego.Router("/user/edit/:id([0-9]+)", &controllers.UserController{}, "POST:Update")

	beego.InsertFilter("/user/delete/:id([0-9]+)", beego.BeforeRouter, filters.HasPermission)
	beego.Router("/user/delete/:id([0-9]+)", &controllers.UserController{}, "GET:Delete")

	beego.Router("/user/:username/topics", &controllers.TopicController{}, "GET:UserTopic")
	beego.Router("/user/:username/replies", &controllers.ReplyController{}, "GET:UserReplay")

	//角色管理
	beego.InsertFilter("/role/list", beego.BeforeRouter, filters.HasPermission)
	beego.Router("/role/list", &controllers.RoleController{}, "GET:List")

	beego.InsertFilter("/role/add", beego.BeforeRouter, filters.HasPermission)
	beego.Router("/role/add", &controllers.RoleController{}, "GET:Add")
	beego.Router("/role/add", &controllers.RoleController{}, "Post:Save")

	beego.InsertFilter("/role/edit/:id([0-9]+)", beego.BeforeRouter, filters.HasPermission)
	beego.Router("/role/edit/:id([0-9]+)", &controllers.RoleController{}, "GET:Edit")
	beego.Router("/role/edit/:id([0-9]+)", &controllers.RoleController{}, "Post:Update")

	beego.InsertFilter("/role/delete/:id([0-9]+)", beego.BeforeRouter, filters.HasPermission)
	beego.Router("/role/delete/:id([0-9]+)", &controllers.RoleController{}, "GET:Delete")

	//权限管理
	beego.InsertFilter("/permission/list", beego.BeforeRouter, filters.HasPermission)
	beego.Router("/permission/list", &controllers.PermissionController{}, "GET:List")
	beego.InsertFilter("/permission/add", beego.BeforeRouter, filters.HasPermission)
	beego.Router("/permission/add", &controllers.PermissionController{}, "GET:Add")
	beego.Router("/permission/add", &controllers.PermissionController{}, "Post:Save")
	beego.InsertFilter("/permission/edit", beego.BeforeRouter, filters.HasPermission)
	beego.Router("/permission/edit/:id([0-9]+)", &controllers.PermissionController{}, "GET:Edit")
	beego.Router("/permission/edit/:id([0-9]+)", &controllers.PermissionController{}, "Post:Update")
	beego.InsertFilter("/permission/delete/:id([0-9]+)", beego.BeforeRouter, filters.HasPermission)
	beego.Router("/permission/delete/:id([0-9]+)", &controllers.PermissionController{}, "GET:Delete")

}
