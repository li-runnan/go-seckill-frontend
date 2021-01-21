package controllers

import (
	"github.com/astaxie/beego/orm"
	"myseckill/models"
)

type EditPersonalInfoController struct {
	BaseController
}

func (e *EditPersonalInfoController) Get(){
	studentNo := e.GetSession("studentNo")
	student := models.Student{Studentid: studentNo.(string)}
	o := orm.NewOrm()
	o.Read(&student,"Studentid")
	e.Data["name"] = student.Name
	e.Data["phone"] = student.Phone
	e.Data["email"] = student.Email
	e.TplName = "EditPersonalInfo.html"
}

func (e *EditPersonalInfoController) Post(){
	studentNo := e.GetSession("studentNo")
	name := e.GetString("name")
	phone := e.GetString("phone")
	email := e.GetString("email")

	student := models.Student{Studentid: studentNo.(string)}
	o := orm.NewOrm()
	o.Read(&student,"Studentid")
	student.Name = name
	student.Phone = phone
	student.Email = email
	o.Update(&student)

	e.SetSession("name",name)

	e.Ctx.ResponseWriter.Write([]byte("<script language='javascript'>alert('保存成功');window.location.href='/edit';</script>"))
}
