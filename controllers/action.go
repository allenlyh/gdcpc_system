package controllers

import (
	. "gdcpc_system/models"
	"github.com/astaxie/beego"
	"strconv"
)

type ActionController struct {
	beego.Controller
}

func (this *ActionController) Get() {
	action := this.GetString("action")
	if action == "DeleteTeam" {
		team_id := this.Input().Get("tid")
		id, _ := strconv.Atoi(team_id)
		DeleteTeam(id)
		this.Redirect("/show_file", 302)
	}
}

func DeleteTeam(id int) {
	var team Teams
	team.Tid = id
	err := team.GetInfoById()
	if err == nil {
		team.Delete()
	}
}
