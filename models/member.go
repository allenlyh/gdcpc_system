package models

import (
	"github.com/astaxie/beego/orm"
)


type Member struct {
	Mid				int `orm:"pk;auto"`
	Username	string
	Password	string
	Ch_name		string
	En_name		string
	Teams		*Teams	`orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Member))
}
