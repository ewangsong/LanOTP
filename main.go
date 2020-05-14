package main

import (
	_ "ewangsong/LanOTP/models"
	_ "ewangsong/LanOTP/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

//FilterUser 这里应该加载 "github.com/astaxie/beego/context" 否则会加载src/context
//过滤器
// var FilterUser = func(ctx *context.Context) {
// 	_, ok := ctx.Input.Session("username").(string)
// 	ok2 := strings.Contains(ctx.Request.RequestURI, "/login")
// 	if !ok && !ok2 {
// 		ctx.Redirect(302, "/login")
// 	}
// }

func init() {
	//beego.AppPath = "/opt/lanradius"
	//beego.SetStaticPath("/static", "/opt/lanradius/static")

	//	beego.LoadAppConfig("ini", "conf/app.conf")

	// jsonConfig := `{
	//     "filename" : "/var/log/lanradius/lanradius.log"
	// }` //定义日志文件路径和名字
	jsonConfig := `{
	    "filename" : "./lanradius.log"
	}` //定义日志文件路径和名字

	logs.SetLogger(logs.AdapterFile, jsonConfig) // 设置日志记录方式：本地文件记录
	logs.EnableFuncCallDepth(true)               // 输出log时能显示输出文件名和行号（非必须）
	//beego.BeeLogger.DelLogger("console")         //删除console日志输出
	//注册过滤器
	//	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	//打开session
	beego.BConfig.WebConfig.Session.SessionOn = true

}

func main() {
	beego.Run()
	//初始化command命令
	//cmd.Cmd()

}
