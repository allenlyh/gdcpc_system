package models

import (
	"github.com/astaxie/beego/orm"
)

type Coach struct {
	Uid      int `orm:"pk;auto"`
	Username string
	Password string
	Chname   string
	Enname   string
	School   string
	Email    string
	Tshirt	string
	Shareroom int
	Accomodate	string
	Admin    int8     `orm:"default(0)"`
	Teams    []*Teams `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Coach))
}

func (this *Coach) GetInfoById() error {
	o := orm.NewOrm()
	err := o.QueryTable("Coach").Filter("uid", this.Uid).One(this)
	return err
}

func (this *Coach) GetInfoByUsername() error {
	o := orm.NewOrm()
	err := o.QueryTable("Coach").Filter("username", this.Username).One(this)
	return err
}

func (this *Coach) CountByUsername() int64 {
	o := orm.NewOrm()
	cnt, _ := o.QueryTable("Coach").Filter("Username", this.Username).Count()
	return cnt
}

func (this *Coach) GetPwdByUsername(username string) (string, error) {
	o := orm.NewOrm()
	var coach Coach
	err := o.QueryTable("Coach").Filter("username", username).One(&coach)
	return coach.Password, err
}

func (this *Coach) Insert() error {
	o := orm.NewOrm()
	_, err := o.Insert(this)
	return err
}

func (this *Coach) Update() error {
	o := orm.NewOrm()
	_, err := o.Update(this)
	return err
}
