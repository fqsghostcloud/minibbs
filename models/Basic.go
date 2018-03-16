package models

import (
	"github.com/astaxie/beego/orm"
)

type Basic struct {
}

var ORM orm.Ormer

func init() {
	ORM = orm.NewOrm()
}
