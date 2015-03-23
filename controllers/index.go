package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	if this.GetSession("logined") != nil {
		this.Data["logined"] = 1
	}
	this.Data["uid"] = this.GetSession("uid")
	this.Data["username"] = this.GetSession("username")
	this.TplNames = "index.tpl"
}
