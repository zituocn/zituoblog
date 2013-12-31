package main

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
	"mime"
	"zituoblog/controllers"
	"zituoblog/models"
)

func init() {
	models.RegisterDB()
}

func main() {
	//orm.Debug = true
	mime.AddExtensionType(".css", "text/css")
	beego.Errorhandler("404", controllers.Page_not_found)
	//前台页面路由
	beego.Router("/", &controllers.HomeHandel{})
	beego.Router("/category/:ename:string", &controllers.CategoryHandel{})
	beego.Router("/category/:ename:string/:page:int", &controllers.CategoryHandel{})
	beego.Router("/view/:id:int", &controllers.ViewHandel{})
	//后台管理路由
	beego.Router("/webadmin", &controllers.AdminLoginHandel{})
	beego.Router("/webadmin/main", &controllers.AdminMainHandel{})
	beego.Router("/webadmin/left", &controllers.AdminLeftHandel{})
	beego.Router("/webadmin/logout", &controllers.AdminLogoutHandel{})

	beego.Router("/webadmin/:object(news)/:action(add|list)", &controllers.AdminNewsHandel{})
	beego.Router("/webadmin/:object(news)/:action(list)/:page([0-9]+)", &controllers.AdminNewsHandel{})
	beego.Router("/webadmin/:object(news)/:action(edit|delete)/:id([0-9]+)", &controllers.AdminNewsHandel{})

	//后台使用的文件上传路由
	beego.Router("/uploadapi/v1", &controllers.UpLoadHandel{})

	//beego.SetStaticPath("/upload", "upload/")

	// //模板函数
	beego.AddFuncMap("navString", models.GetNavString) //返回所有页面的动态导航栏
	beego.HttpPort = 33638
	beego.RunMode = "pro"
	beego.Run()
}
