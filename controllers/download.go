package controllers

import (
	"github.com/astaxie/beego"
)

type DownloadController struct {
	beego.Controller
}

func (this *DownloadController) Get() {
	this.TplNames = "download.tpl"
}
