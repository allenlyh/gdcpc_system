package models

import (
	"github.com/astaxie/beego/orm"
)


type Teams struct {
	Tid				int `orm:"pk;auto"`
	Ch_name		string
	En_name		string
	Coach		*Coach	`orm:"rel(fk)"`
	Member		[]*Member	`orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Teams))
}
