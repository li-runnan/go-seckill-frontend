package controllers

import "github.com/astaxie/beego"

type LogoutController struct {
	beego.Controller
}

func (l *LogoutController) Get() {
	l.DestroySession()
	l.Redirect("/login",302)
}
