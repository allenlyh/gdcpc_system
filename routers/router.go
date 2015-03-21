package routers

import (
	"gdcpc_system/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.IndexController{})
    beego.Router("/index", &controllers.IndexController{})
}
