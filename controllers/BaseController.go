package controllers

import (
	"github.com/astaxie/beego"
)

//等同于php 控制器中的common.php
type BaseController struct {
	beego.Controller
}

//重写Prepare方法,通过session判断用户是否登陆，若未登陆则跳转到登陆页面
func (b *BaseController) Prepare(){
	studentNo := b.GetSession("studentNo")
	name := b.GetSession("name")
	//用户未登录，跳转至登录界面
	if studentNo == nil {
		b.Ctx.ResponseWriter.Write([]byte("<script language='javascript'>alert('您还未登录，请先登录');window.location.href='/login';</script>"))
	}
	b.Data["studentNo"] = studentNo
	b.Data["name"] = name
}

//获取用户权限
func (b *BaseController) GetAuthor(){

}
