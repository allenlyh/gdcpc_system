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
		team  Teams
		err   error
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
	team.Tid, err = strconv.Atoi(this.GetString("tid"))
	coach.Uid = this.GetSession("uid").(int)
	coach.GetInfoById()
	if err == nil && team.Tid != 0 {
		team.GetInfoById()
		if team.Coach.Uid != coach.Uid {
			this.Data["warning"] = "Illegel operation!"
			this.TplNames = "warning.tpl"
			return
		}
		this.Data["init"] = team
	//} else {
	//	this.Data["warning"] = "You can't create team now."
	//	this.TplNames = "warning.tpl"
	//	return
	}
	this.TplNames = "create_team.tpl"
}

func (this *CreateTeamController) Post() {
	var (
		coach     Coach
		team      Teams
		check_err error
		err       error
		flag      int
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
	flag = 0
	team.Tid, err = strconv.Atoi(this.GetString("tid"))
	if err == nil && team.Tid != 0 {
		team.GetInfoById()
		if team.Coach.Uid != coach.Uid {
			this.Data["warning"] = "Illegel operation!"
			this.TplNames = "warning.tpl"
			return
		}
	}

	team.Region, check_err = CheckNotEmpty(this.GetString("region"))
	if check_err != nil {
		this.Data["warning"] = check_err
		flag = 1
	}

	team.Ch_name, check_err = CheckNotEmpty(this.GetString("chinese_name"))
	if check_err != nil {
		this.Data["warning"] = check_err
		flag = 1
	}
	team.En_name, check_err = CheckEnglish(this.GetString("english_name"))
	if check_err != nil {
		this.Data["warning"] = check_err
		flag = 1
	}
	team.Mem1chname, check_err = CheckNotEmpty(this.GetString("mem1_chname"))
	if check_err != nil {
		this.Data["warning"] = check_err
		flag = 1
	}
	team.Mem1enname, check_err = CheckPT(this.GetString("mem1_enname"))
	if check_err != nil {
		this.Data["warning"] = check_err
		flag = 1
	}
	team.Mem1email, check_err = CheckEmail(this.GetString("mem1_email"))
	if check_err != nil {
		this.Data["warning"] = check_err
		flag = 1
	}
	team.Sex1, check_err = strconv.Atoi(this.GetString("sex1"))
	team.Mem2chname, check_err = CheckNotEmpty(this.GetString("mem2_chname"))
	if check_err != nil {
		this.Data["warning"] = check_err
		flag = 1
	}
	team.Mem2enname, check_err = CheckPT(this.GetString("mem2_enname"))
	if check_err != nil {
		this.Data["warning"] = check_err
		flag = 1
	}
	team.Mem2email, check_err = CheckEmail(this.GetString("mem2_email"))
	if check_err != nil {
		this.Data["warning"] = check_err
		flag = 1
	}
	team.Sex2, check_err = strconv.Atoi(this.GetString("sex2"))
	team.Mem3chname, check_err = CheckNotEmpty(this.GetString("mem3_chname"))
	if check_err != nil {
		this.Data["warning"] = check_err
		flag = 1
	}
	team.Mem3enname, check_err = CheckPT(this.GetString("mem3_enname"))
	if check_err != nil {
		this.Data["warning"] = check_err
		flag = 1
	}
	team.Mem3email, check_err = CheckEmail(this.GetString("mem3_email"))
	if check_err != nil {
		this.Data["warning"] = check_err
		flag = 1
	}
	team.Sex3, check_err = strconv.Atoi(this.GetString("sex3"))
	team.Tshirt, check_err = CheckNotEmpty(this.GetString("tshirt"))
	if check_err != nil {
		this.Data["warning"] = check_err
		flag = 1
	}
	team.Coach = &coach
	team.School = coach.School
	team.Coachnamech = coach.Chname
	team.Coachnameen = coach.Enname
	team.Coachemail = coach.Email
	fmt.Println(err)
	if flag == 1 {
		this.Data["init"] = team
		this.TplNames = "create_team.tpl"
		return
	}
	if err == nil && team.Tid != 0 {
		team.Update()
	} else {
		team.Insert()
	}
	this.Redirect("/show_file", 302)
}
