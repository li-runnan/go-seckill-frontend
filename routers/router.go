//router.go

// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"myseckill/controllers"
)

func init() {
	//add route
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
		beego.NSNamespace("/home",
			beego.NSInclude(
				&controllers.HomeController{},
			),
		),
	)
	beego.AddNamespace(ns)
	beego.Router("/",&controllers.LoginController{})
	beego.Router("/login",&controllers.LoginController{})
	beego.Router("/home",&controllers.HomeController{})
	beego.Router("/welcome",&controllers.WelcomeController{})
	beego.Router("/logout",&controllers.LogoutController{})
	beego.Router("/edit",&controllers.EditPersonalInfoController{})
	beego.Router("/result",&controllers.SecKillResultController{})
	beego.Router("/seckill",&controllers.SecKillController{})
}
