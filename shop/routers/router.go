package routers

import (
	"mygo/shop/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/index", &controllers.MainController{})
}
