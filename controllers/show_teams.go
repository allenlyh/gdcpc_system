package controllers

import (
	. "gdcpc_system/models"
	"github.com/astaxie/beego"
)

type ShowTeamsController struct {
	beego.Controller
}

func (this *ShowTeamsController) Get() {
	var (
		coach Coach
		err   error
		team  Teams
	)
	if this.GetSession("logined") == nil {
		this.Data["warning"] = "Please login!"
		this.TplNames = "warning.tpl"
		return
	} else {
		this.Data["logined"] = 1
	}
	coach.Uid = this.GetSession("uid").(int)
	err = coach.GetInfoById()
	if err != nil {
		this.Data["warning"] = "No such user!"
		this.TplNames = "warning.tpl"
		return
	}
	if coach.Admin == 0 {
		this.Data["uid"] = this.GetSession("uid")
		this.Data["username"] = this.GetSession("username")
		this.Data["warning"] = "Illegal operation!"
		this.TplNames = "warning.tpl"
		return
	}
	this.Data["uid"] = this.GetSession("uid")
	this.Data["username"] = this.GetSession("username")
	this.Data["total_teams"], _ = team.GetAllTeams()
	this.TplNames = "show_teams.tpl"
}
