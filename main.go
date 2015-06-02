package main

import (
	_ "gdcpc_system/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	kDataBaseUser     = "gdcpc"
	kDataBasePassword = "gdcpcadmin"
	kDataBaseName     = "gdcpc"
)

func init() {
	orm.Debug = true
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", kDataBaseUser+":"+kDataBasePassword+"@/"+kDataBaseName+"?charset=utf8&loc=Asia%2FShanghai")
}

func main() {
	orm.RunCommand()
	beego.SetStaticPath("/download", "download")
	beego.Run()
}
