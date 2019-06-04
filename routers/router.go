package routers

import (
	"github.com/astaxie/beego"
	"network/controllers"
)

func init() {
	//beego.Router("/", &controllers.MainController{})

	ns := beego.NewNamespace("/sl",
		beego.NSNamespace(
			"/v1",
			beego.NSRouter("/users", &controllers.UserController{}),
		),
	)
	beego.AddNamespace(ns)

}
