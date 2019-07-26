package routers

import (
	"XingYiYun/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login",&controllers.LoginController{})
    beego.Router("/welcome",&controllers.WelController{})
    beego.Router("/banner",&controllers.BannerController{})
    beego.Router("/addBanner", &controllers.AddBanner{})
    beego.Router("/user",&controllers.UserController{})
    beego.Router("/about",&controllers.AboutController{})
}
