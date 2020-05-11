package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func BsaDelete(id int) {
	o := orm.NewOrm()
	bas := WsBas{Id: id}

	_, err := o.Delete(&bas)
	if err != nil {
		beego.Info("删除bas错误", err)
		return
	}
}

func UserDelete(id int) {
	o := orm.NewOrm()
	user := WsUsers{Id: id}

	_, err := o.Delete(&user)
	if err != nil {
		beego.Info("删除用户错误", err)
		return
	}
}
