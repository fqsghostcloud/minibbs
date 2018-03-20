package main

import (
	"minibbs/models"
	_ "minibbs/routers"
	_ "minibbs/templates"
	_ "minibbs/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("jdbc.username")+":"+beego.AppConfig.String("jdbc.password")+"@tcp(192.168.34.145)/pybbsgo?charset=utf8&parseTime=true&charset=utf8&loc=Asia%2FShanghai", 30)
	orm.RegisterModel(
		new(models.User),
		new(models.Topic),
		new(models.Section),
		new(models.Reply),
		new(models.ReplyUpLog),
		new(models.Role),
		new(models.Permission))
	orm.RunSyncdb("default", false, true)
}

func main() {
	//orm.Debug = true
	//ok, err := regexp.MatchString("/topic/edit/[0-9]+", "/topic/edit/123")
	//beego.Debug(ok, err)

	// use glog----------------------------
	// flag.Parse()
	// defer glog.Flush()

	// orm.Debug = true                                 // database debug model
	beego.BConfig.WebConfig.Session.SessionOn = true // session on
	beego.Run()
	// glog.Flush()
}
