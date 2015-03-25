package models

import (
	"github.com/astaxie/beego/orm"
)

type Teams struct {
	Tid         int `orm:"pk;auto"`
	Ch_name     string
	En_name     string
	Coach       *Coach `orm:"rel(fk)"`
	Region      string
	Coachnamech string
	Coachnameen string
	School      string
	Mem1chname  string
	Mem1enname  string
	Mem1email   string
	Sex1        int
	Mem2chname  string
	Mem2enname  string
	Mem2email   string
	Sex2        int
	Mem3chname  string
	Mem3enname  string
	Mem3email   string
	Sex3        int
}

func init() {
	orm.RegisterModel(new(Teams))
}

func (this *Teams) GetTeamsByCoach(region string) ([]Teams, error) {
	var (
		all []Teams
	)
	o := orm.NewOrm()
	_, err := o.QueryTable("teams").Filter("Coach", this.Coach).Filter("Region", region).All(&all)
	return all, err
}

func (this *Teams) GetAllTeams(region string) ([]Teams, error) {
	var (
		all []Teams
	)
	o := orm.NewOrm()
	_, err := o.QueryTable("teams").Filter("Region", region).OrderBy("School").All(&all)
	return all, err
}

func (this *Teams) GetInfoById() error {
	o := orm.NewOrm()
	err := o.QueryTable("Teams").Filter("tid", this.Tid).One(this)
	return err
}

func (this *Teams) Insert() error {
	o := orm.NewOrm()
	_, err := o.Insert(this)
	return err
}

func (this *Teams) Update() error {
	o := orm.NewOrm()
	_, err := o.Update(this)
	return err
}

func (this *Teams) Delete() error {
	o := orm.NewOrm()
	_, err := o.Delete(this)
	return err
}
