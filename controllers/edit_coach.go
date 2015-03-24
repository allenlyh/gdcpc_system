package controllers

import (
	"fmt"
	. "gdcpc_system/models"
	. "gdcpc_system/tools"
	"github.com/astaxie/beego"
)

type EditCoachController struct {
	beego.Controller
}

func (this *EditCoachController) Get() {
	var (
		coach Coach
	)
	if this.GetSession("logined") != nil {
		this.Data["logined"] = 1
	} else {
		this.Data["warning"] = "Please login!"
		this.TplNames = "warning.tpl"
		return
	}
	this.Data["uid"] = this.GetSession("uid")
	this.Data["username"] = this.GetSession("username")
	coach.Uid = this.GetSession("uid").(int)
	coach.GetInfoById()
	fmt.Println(coach.Name)
	this.Data["init"] = coach
	this.TplNames = "edit_coach.tpl"
}

func (this *EditCoachController) Post() {
	var (
		coach     Coach
		check_err error
		_         error
		old_pwd   string
		new_pwd1  string
		new_pwd2  string
	)
	if this.GetSession("logined") != nil {
		this.Data["logined"] = 1
	} else {
		this.Data["warning"] = "Please login!"
		this.TplNames = "warning.tpl"
		return
	}
	this.Data["uid"] = this.GetSession("uid")
	this.Data["username"] = this.GetSession("username")
	coach.Uid = this.GetSession("uid").(int)
	coach.GetInfoById()
	this.Data["init"] = coach
	old_pwd, check_err = GetMD5Pwd(this.GetString("old_password"))
	if old_pwd != coach.Password {
		this.Data["warning"] = "Wrong old password!"
		this.TplNames = "edit_coach.tpl"
		return
	}
	new_pwd1 = this.GetString("new_password1")
	if len(new_pwd1) != 0 {
		new_pwd1, check_err = GetMD5Pwd(new_pwd1)
		new_pwd2, check_err = GetMD5Pwd(this.GetString("new_password2"))
		if new_pwd2 != new_pwd1 {
			this.Data["warning"] = "Your passwords do not match!"
			this.TplNames = "edit_coach.tpl"
			return
		}
		if check_err != nil {
			this.Data["warning"] = check_err
			this.TplNames = "edit_coach.tpl"
			return
		}
		coach.Password = new_pwd1
	}
	coach.Name, check_err = CheckNotEmpty(this.GetString("name"))
	if check_err != nil {
		this.Data["warning"] = "Name can't be empty!"
		this.TplNames = "edit_coach.tpl"
		return
	}
	coach.School, check_err = CheckNotEmpty(this.GetString("school"))
	if check_err != nil {
		this.Data["warning"] = "School can't be empty!"
		this.TplNames = "edit_coach.tpl"
		return
	}
	coach.Update()
	this.Redirect("/show_file", 302)
}