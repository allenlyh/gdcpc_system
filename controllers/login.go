package controllers

import (
	. "gdcpc_system/models"
	. "gdcpc_system/tools"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplNames = "login.tpl"
}

func (this *LoginController) Post() {
	var (
		password string
		coach    Coach
		_        error
	)
	coach.Username = this.GetString("username")
	if coach.CountByUsername() == 0 {
		this.Data["warning"] = "Username or Password wrong!"
		this.TplNames = "login.tpl"
		return
	}
	coach.GetInfoByUsername()
	password = this.GetString("password")
	password, _ = GetMD5Pwd(password)
	if coach.Password != password {
		this.Data["warning"] = "Username or Password wrong!"
		this.TplNames = "login.tpl"
		return
	}
	this.SetSession("logined", 1)
	this.SetSession("uid", coach.Uid)
	this.SetSession("username", coach.Username)
	this.Redirect("/index", 302)
}
