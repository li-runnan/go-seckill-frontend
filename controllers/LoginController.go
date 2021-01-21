package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myseckill/models"
	"time"
)

type LoginController struct {
	beego.Controller
}

// @Title Login
// @Description login
// @router / [get]
func (l *LoginController) Get(){
	l.TplName = "login.html"
}

func (l *LoginController) Post(){
	username := l.GetString("username")
	password := md5.Sum([]byte(l.GetString("password")))
	md5password := hex.EncodeToString(password[:])
	fmt.Println("用户名:",username,"密码:",md5password)

	student := models.Student{Studentid: username}
	o := orm.NewOrm()
	err := o.Read(&student, "Studentid")
	fmt.Println(student)

	if err == orm.ErrNoRows || student.Password != md5password {
		l.Ctx.ResponseWriter.Write([]byte("<script language='javascript'>alert('用户名或密码错误');window.location.href='/login';</script>"))
	} else {
		l.SetSession("studentNo",username)
		l.SetSession("name",student.Name)
		student.Ipaddress = l.GetClientIP()
		student.Lastlogintime = time.Now()
		o.Update(&student)
		l.Redirect("/home",302)
	}
}

//获取用户IP地址
func (l *LoginController) GetClientIP() string{
	s := l.Ctx.Input.IP()
	if s == "::1"{
		s = "127.0.0.1"
	}
	return s
}

