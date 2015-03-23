package controllers

import (
	. "gdcpc_system/models"
	"github.com/astaxie/beego"
)

type ShowFileController struct {
	beego.Controller
}

func (this *ShowFileController) Get() {
	var (
		coach Coach
		err   error
		team  Teams
	)
	if this.GetSession("logined") == nil {
		this.Data["warning"] = "Please login!"
		this.TplNames = "show_file.tpl"
		return
	}
	coach.Uid = this.GetSession("uid").(int)
	err = coach.GetInfoById()
	if err != nil {
		this.Data["warning"] = "No such user!"
	} else {
		this.Data["show"] = coach
	}
	if this.GetSession("logined") != nil {
		this.Data["logined"] = 1
	}
	this.Data["uid"] = this.GetSession("uid")
	this.Data["username"] = this.GetSession("username")
	team.Coach = &coach
	this.Data["total_teams"], _ = team.GetTeamsByCoach()
	this.Data["coachname"] = coach.Name
	this.TplNames = "show_file.tpl"
}
