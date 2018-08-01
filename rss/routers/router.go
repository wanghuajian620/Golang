package routers

import (
	"github.com/wanghuajian620/Golang/rss/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
