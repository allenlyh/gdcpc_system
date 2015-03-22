package controllers

import (
	"github.com/astaxie/beego"
	. "gdcpc_system/models"
	"tools"
	"strconv"
)

type CoachController struct {
	beego.Controller
}

func (this *CoachController) Get() {
	var (
		err error
		id string
		coach Coach
	)
	id = this.GetString("uid")
	table.Uid, err = strconv.Atoi(id)
	if err == nil && table.Uid != 0 {
		err = coach.GetInfoById()
		this.Data["init"] = coach
	}
	this.TplNames = "edit_coach.tpl"
}

func (this *CoachController) Post() {
	var (
		coach Coach
		check_err error
		err error
	)
	coach.Uid, err = strconv.Atoi(this.GetString("uid"))
	err = coach.GetInfoById()
	coach.Username, check_err = CheckUserName(this.GetString("username"))
	if check_err != nil {
		this.Data["warning"] = check_err
		this.TplNames = "edit_coach.tpl"
		return
	}
	coach.Password, check_err = GetMD5Pwd(this.GetString("password"))
	if check_err != nil {
		this.Data["warning"] = check_err
		this.TplNames = "edit_coach.tpl"
	}
	coach.NickName, check_err = CheckNotEmpty(this.GetString("nickname"))
	if check_err != nil {
		this.Data["warning"] = "Nickname can't be empty!"
		this.TplNames = "edit_coach.tpl"
	}
	coach.School, check_err = CheckNotEmpty(this.GetString("School"))
	if check_err != nil {
		this.Data["warning"] = "School can't be empty!"
		this.TplNames = "edit_coach.tpl"
	}
	coack.Email, check_err = CheckEmail(this.GetString("email"))
	if check_err != nil {
		this.Data["warning"] = check_err
		this.TplNames = "edit_coach.tpl"
	} else {
		if err == nil && table.Uid != 0 {
			coach.Update()
		} else {
			coach.Insert()
			this.Data["warning"] = "You have succeeded in registering! Login please!"
		}
		this.Redirect("/index", 302)
	}
}
