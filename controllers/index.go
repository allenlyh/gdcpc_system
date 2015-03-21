package controllers

import (
	"github.com/astaxie/beego"
	. "gdcpc_system/models"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	var (
		coach Coach
	)
	this.Data["Post"] = coach
	this.TplNames = "index.tpl"
}
