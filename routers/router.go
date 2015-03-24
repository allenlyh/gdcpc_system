package routers

import (
	"gdcpc_system/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/reg_coach", &controllers.RegCoachController{})
	beego.Router("/edit_coach", &controllers.EditCoachController{})
	beego.Router("/create_team", &controllers.CreateTeamController{})
	beego.Router("/show_team", &controllers.ShowTeamsController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/show_file", &controllers.ShowFileController{})
	beego.Router("/action", &controllers.ActionController{})
}
