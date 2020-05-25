package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func BsaDelete(id int) WsBas {
	o := orm.NewOrm()
	bas := WsBas{Id: id}
	o.Read(&bas)
	_, err := o.Delete(&bas)
	if err != nil {
		beego.Info("删除bas错误", err)
	}
	return bas
}

//UserDelete 删除用户
func UserDelete(id int) WsUsers {
	o := orm.NewOrm()
	user := WsUsers{Id: id}

	_, err := o.Delete(&user)
	if err != nil {
		beego.Info("删除用户错误", err)
	}
	return user
}

//TokenDelete 删除token
func TokenDelete(id int) WsOtp {
	o := orm.NewOrm()
	token := WsOtp{Id: id}
	o.Read(&token)
	_, err := o.Delete(&token)
	if err != nil {
		beego.Info("删除token错误", err)
	}
	return token
}
