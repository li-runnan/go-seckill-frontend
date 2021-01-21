package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Tickets struct {
	Id int `orm:"pk"`
	Name string
	Num int
	Starttime time.Time
}

func GetTicketsInfoById(id int) Tickets {
	var ticket Tickets
	o := orm.NewOrm()
	o.QueryTable("Tickets").Filter("id",id).All(&ticket)
	return ticket
}

func GetTicketsInfo() []Tickets {
	var ticket []Tickets
	o := orm.NewOrm()
	o.QueryTable("Tickets").All(&ticket)
	return ticket
}

func GetTicketsInfoIn24hours() []Tickets {
	var ticket []Tickets
	var res []Tickets
	o := orm.NewOrm()
	o.QueryTable("Tickets").All(&ticket)
	count := 0
	for i := 0; i< len(ticket);i++ {
		t1 := ticket[i].Starttime
		t2 := time.Now()
		interval := t1.Sub(t2)
		if interval.Hours() < 24 {
			res = append(res,ticket[i])
			count = count + 1
		}
	}
	return res
}

