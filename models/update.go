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
			return "no"
		}
		return "ok"
	} else {
		return "no"
	}

}

//bas修改
func BasUpdate(id int, name, ip_add, secret, port string) (oldbas, newbas WsBas) {
	oldbas = WsBas{Id: id}
	o := orm.NewOrm()
	err1 := o.Read(&oldbas)
	if err1 != nil {
		beego.Info("bas更新错误", err1)
	}
	newbas = oldbas
	newbas.Name = name
	newbas.IpAddr = ip_add
	newbas.Secret = secret
	newbas.Port = port

	_, err := o.Update(&newbas)

	if err != nil {
		beego.Info("bas更新错误", err)
	}

	return oldbas, newbas

}

//用户修改
func UserUdate(id int, realname, name string) (olduser, newuser WsUsers) {
	olduser = WsUsers{Id: id}
	o := orm.NewOrm()
	err1 := o.Read(&olduser)
	newuser = olduser
	if err1 != nil {
		beego.Info("用户更新错误", err1)
	}
	newuser.Name = name
	newuser.RealName = realname
	_, err := o.Update(&newuser)

	if err != nil {
		beego.Info("用户更新错误", err)
	}
	return olduser, newuser
}

//token 修改
func TokenUdate(id int, name string) WsOtp {
	token := WsOtp{Id: id}
	o := orm.NewOrm()
	err1 := o.Read(&token)
	if err1 != nil {
		beego.Info("用户更新查询错误", err1)
	}
	token.BindingUser = name
	_, err := o.Update(&token)

	if err != nil {
		beego.Info("token更新错误", err)
	}
	return token
}
