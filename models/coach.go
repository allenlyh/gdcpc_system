package models

import (
	"github.com/astaxie/beego/orm"
)


type Coach struct {
	Username		string `orm:"pk"`
	Password	string
	Name	string	
	School			string
	Email	string
	Teams	[]*Teams `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Coach))
}

func (this *Coach) GetById(id int) error {
	o := orm.NewOrm()
	err := o.QueryTable("Coach").Filter("uid", id).One(this)
	return err;
}

func (this *Coach) GetPwdByUsername(username string) (string, error) {
	o := orm.NewOrm()
	var coach Coach
	err := o.QueryTable("Coach").Filter("username", username).One(&coach)
	return coach.Password, err
}

func (this *Coach) Register() error {
	o := orm.NewOrm()
	_, err := o.Insert(this)
	return err
}

func (this *Coach) Update() error {
	o := orm.NewOrm()
	_, err := o.Update(this)
	return err
}
