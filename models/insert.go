package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func BasInsert(name, ip_add, secret, port string) bool {
	bas := WsBas{}

	o := orm.NewOrm()
	bas.Name = name
	bas.IpAddr = ip_add
	bas.Secret = secret
	bas.Port = port
	_, err := o.Insert(&bas)
	if err != nil {
		beego.Info("bas插入错误", err)
		return false
	} else {
		return true
	}

}

//添加用户数据库操作

func UserInsert(realname, name, password string) bool {
	user := WsUsers{RealName: realname, Name: name, Password: password}
	o := orm.NewOrm()
	if realname == "" || name == "" || password == "" {
		return false
	}

	err := o.Read(&user, "Name") //先查询是否存在此用户

	if err != nil {
		_, err := o.Insert(&user)
		if err != nil {
			beego.Info("用户插入错误", err)
			return false
		}
	} else {
		return false
	}
	return true
}

//添加默认管理员账号密码
func AddAdmin() {
	admin := WsAdmin{Name: "admin", Password: "admin"}
	o := orm.NewOrm()
	err := o.Read(&admin, "Name") //先查询是否存在admin管理员

	if err != nil { //有错不存在管理员
		_, err := o.Insert(&admin) //插入默认管理员
		if err != nil {            //插入有错误退出程序
			beego.Info("添加默认管理员插入错误", err)
			return
		}
	}
}

//插入操作日志
func LogInsert() {

}
