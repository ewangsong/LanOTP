package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//管理admin更新
func AdminUpdate(a, b string) string {
	o := orm.NewOrm()
	admin := WsAdmin{}
	if a != b {
		return "on"
	}

	admin.Name = "admin"
	if o.Read(&admin, "Name") == nil {
		admin.Password = a
		_, err := o.Update(&admin)
		if err != nil {
			return "on"
		}
		return "ok"
	} else {
		return "on"
	}

}

//bas修改
func BasUpdate(id int, name, ip_add, secret, port string) {
	bas := WsBas{Id: id}
	o := orm.NewOrm()
	err1 := o.Read(&bas)
	if err1 != nil {
		beego.Info("bas更新错误", err1)
	}
	bas.Name = name
	bas.IpAddr = ip_add
	bas.Secret = secret
	bas.Port = port

	_, err := o.Update(&bas)

	if err != nil {
		beego.Info("bas更新错误", err)
	}

}

//用户修改
func UserUdate(id int, realname, name, password string) {
	user := WsUsers{Id: id}
	o := orm.NewOrm()
	err1 := o.Read(&user)
	if err1 != nil {
		beego.Info("用户更新错误", err1)
	}
	user.Name = name
	user.RealName = realname
	user.Password = password
	_, err := o.Update(&user)

	if err != nil {
		beego.Info("用户更新错误", err)
	}

}
