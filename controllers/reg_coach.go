package controllers

import (
	. "gdcpc_system/models"
	. "gdcpc_system/tools"
	"github.com/astaxie/beego"
)

type RegCoachController struct {
	beego.Controller
}

func (this *RegCoachController) Get() {
	this.TplNames = "reg_coach.tpl"
}

func (this *RegCoachController) Post() {
	var (
		coach     Coach
		check_err error
		_         error
		pwd2      string
	)
	coach.Username, check_err = CheckUserName(this.GetString("username"))
	if check_err != nil {
		this.Data["warning"] = check_err
		this.TplNames = "reg_coach.tpl"
		return
	}
	cnt := coach.CountByUsername()
	if cnt != 0 {
		this.Data["warning"] = "The username has been used!"
		this.TplNames = "reg_coach.tpl"
		return
	}
	coach.Password, check_err = GetMD5Pwd(this.GetString("password"))
	if check_err != nil {
		this.Data["warning"] = check_err
		this.TplNames = "reg_coach.tpl"
		return
	}
	pwd2, check_err = GetMD5Pwd(this.GetString("password2"))
	if pwd2 != coach.Password {
		this.Data["warning"] = "Your passwords do not match!"
		this.TplNames = "reg_coach.tpl"
		return
	}
	coach.Name, check_err = CheckNotEmpty(this.GetString("name"))
	if check_err != nil {
		this.Data["warning"] = "Name can't be empty!"
		this.TplNames = "reg_coach.tpl"
		return
	}
	coach.Email, check_err = CheckEmail(this.GetString("email"))
	if check_err != nil {
		this.Data["warning"] = check_err
		this.TplNames = "reg_coach.tpl"
		return
	}
	coach.School, check_err = CheckNotEmpty(this.GetString("school"))
	if check_err != nil {
		this.Data["warning"] = "School can't be empty!"
		this.TplNames = "reg_coach.tpl"
		return
	}
	coach.Insert()
	this.SetSession("logined", 1)
	this.SetSession("uid", coach.Uid)
	this.SetSession("username", coach.Username)
	this.Redirect("/index", 302)
}
