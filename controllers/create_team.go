package controllers

import (
	"fmt"
	. "gdcpc_system/models"
	. "gdcpc_system/tools"
	"github.com/astaxie/beego"
	"strconv"
)

type CreateTeamController struct {
	beego.Controller
}

func (this *CreateTeamController) Get() {
	var (
		team Teams
		err  error
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
	team.Tid, err = strconv.Atoi(this.GetString("tid"))
	if err == nil && team.Tid != 0 {
		team.GetInfoById()
		this.Data["init"] = team
	}
	this.TplNames = "create_team.tpl"
}

func (this *CreateTeamController) Post() {
	var (
		coach     Coach
		team      Teams
		check_err error
		err       error
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
	team.Tid, err = strconv.Atoi(this.GetString("tid"))

	if err == nil && team.Tid != 0 {
		team.GetInfoById()
		this.Data["init"] = team
	}

	team.Ch_name, check_err = CheckNotEmpty(this.GetString("chinese_name"))
	if check_err != nil {
		this.Data["warning"] = check_err
		this.TplNames = "create_team.tpl"
		return
	}
	team.En_name, check_err = CheckEnglish(this.GetString("english_name"))
	if check_err != nil {
		this.Data["warning"] = check_err
		this.TplNames = "create_team.tpl"
		return
	}
	team.Mem1chname, check_err = CheckNotEmpty(this.GetString("mem1_chname"))
	if check_err != nil {
		this.Data["warning"] = check_err
		this.TplNames = "create_team.tpl"
		return
	}
	team.Mem1enname, check_err = CheckPT(this.GetString("mem1_enname"))
	if check_err != nil {
		this.Data["warning"] = check_err
		this.TplNames = "create_team.tpl"
		return
	}
	team.Mem1email, check_err = CheckEmail(this.GetString("mem1_email"))
	if check_err != nil {
		this.Data["warning"] = check_err
		this.TplNames = "create_team.tpl"
		return
	}
	team.Sex1, check_err = strconv.Atoi(this.GetString("sex1"))
	team.Mem2chname, check_err = CheckNotEmpty(this.GetString("mem2_chname"))
	if check_err != nil {
		this.Data["warning"] = check_err
		this.TplNames = "create_team.tpl"
		return
	}
	team.Mem2enname, check_err = CheckPT(this.GetString("mem2_enname"))
	if check_err != nil {
		this.Data["warning"] = check_err
		this.TplNames = "create_team.tpl"
		return
	}
	team.Mem2email, check_err = CheckEmail(this.GetString("mem2_email"))
	if check_err != nil {
		this.Data["warning"] = check_err
		this.TplNames = "create_team.tpl"
		return
	}
	team.Sex2, check_err = strconv.Atoi(this.GetString("sex2"))
	team.Mem3chname, check_err = CheckNotEmpty(this.GetString("mem3_chname"))
	if check_err != nil {
		this.Data["warning"] = check_err
		this.TplNames = "create_team.tpl"
		return
	}
	team.Mem3enname, check_err = CheckPT(this.GetString("mem3_enname"))
	if check_err != nil {
		this.Data["warning"] = check_err
		this.TplNames = "create_team.tpl"
		return
	}
	team.Mem3email, check_err = CheckEmail(this.GetString("mem3_email"))
	if check_err != nil {
		this.Data["warning"] = check_err
		this.TplNames = "create_team.tpl"
		return
	}
	team.Sex3, check_err = strconv.Atoi(this.GetString("sex3"))
	team.Coach = &coach
	team.Coachname = coach.Name
	team.School = coach.School
	fmt.Println(err)
	if err == nil && team.Tid != 0 {
		team.Update()
	} else {
		team.Insert()
	}
	this.Redirect("/show_file", 302)
}
