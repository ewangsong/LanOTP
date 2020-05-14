package routers

import (
	"ewangsong/LanOTP/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &controllers.MainController{}, "get:Getindex")
	beego.Router("/register", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{}, "get:Login;post:PostLogin")
	beego.Router("/about", &controllers.MainController{}, "get:ShowAbout")

	//bas路由
	beego.Router("/admin/bas", &controllers.BasController{}, "get:Getbas")
	beego.Router("/admin/bas/add", &controllers.BasController{}, "get:AddBas;post:PostAddBas")
	beego.Router("/admin/bas/update", &controllers.BasController{}, "get:UpdateBas;post:PostUpdateBas")
	beego.Router("/admin/bas/delete", &controllers.BasController{}, "get:DeleteBas")

	//	beego.Router("/admin/config", &controllers.MainController{}, "get:ShowConfig")
	beego.Router("/admin/superrpc", &controllers.MainController{}, "get:ShowSuperrpc")
	beego.Router("/admin/dashboard", &controllers.MainController{}, "get:ShowDashboard")
	beego.Router("/admin/password", &controllers.MainController{}, "get:ShowChangePassword;post:PostChangePassword")
	beego.Router("/admin/logout", &controllers.MainController{}, "get:LogOut")
	beego.Router("/admin/log", &controllers.MainController{}, "get:ShowLog")

	//user路由
	beego.Router("/admin/users", &controllers.UserController{}, "get:ShowUsers;post:PostShowUsers")
	beego.Router("/admin/users/detail", &controllers.UserController{}, "get:DetailUsers")
	beego.Router("/admin/users/update", &controllers.UserController{}, "get:UpdateUsers;post:PostUpdateUsers")
	beego.Router("/admin/users/delete", &controllers.UserController{}, "get:DeleteUser")
	beego.Router("/admin/users/add", &controllers.UserController{}, "get:AddUser;post:PostAddUser")
	//token 路由
	beego.Router("/admin/token", &controllers.TokenController{}, "get:ShowToken;post:PostShowToken")
	beego.Router("/admin/token/detail", &controllers.TokenController{}, "get:DetailToken")
	beego.Router("/admin/token/update", &controllers.TokenController{}, "get:UpdateToken;post:PostUpdateToken")
	beego.Router("/admin/token/delete", &controllers.TokenController{}, "get:DeleteToken")
	beego.Router("/admin/token/add", &controllers.TokenController{}, "get:AddToken;post:PostAddToken")
}
