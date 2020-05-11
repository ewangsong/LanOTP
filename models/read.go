package models

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func BasRead() []WsBas {
	var BasShow []WsBas
	bas := WsBas{}
	o := orm.NewOrm()
	qs := o.QueryTable(bas)
	if count, err := qs.All(&BasShow); err == nil {
		beego.Info(count, BasShow)
	} else {
		beego.Info("查询失败")
	}

	return BasShow
}

func BasRead2(id int) []WsBas {
	o := orm.NewOrm()
	bas := WsBas{}
	bas.Id = id
	err := o.Read(&bas)
	if err != nil {
		beego.Info("查询bas错误", err)
	}
	var bass []WsBas
	bass = append(bass, bas)
	return bass
}

func UsersRead() (UsersShow []WsUsers, co int) {
	//	var  []WsUsers
	user := WsUsers{}
	o := orm.NewOrm()
	qs := o.QueryTable(user)
	if count, err := qs.All(&UsersShow); err == nil {
		beego.Info(count, UsersShow)
		co = int(count)
	} else {
		beego.Info("查询失败")
	}
	//	PageSize := 2
	//	PageCount := int(math.Ceil((float64(count) / float64(PageSize))))

	return UsersShow, co
}

func UsersRead2(id int) []WsUsers {
	user := WsUsers{Id: id}
	o := orm.NewOrm()
	err := o.Read(&user)
	if err != nil {
		beego.Info("查询用户错误", err)
	}
	var userss []WsUsers
	userss = append(userss, user)
	return userss
}

//用户查询页函数以账户进行查询方法
func UsersRead3(s string) []WsUsers {
	var UsersShow []WsUsers
	user := WsUsers{}
	o := orm.NewOrm()
	if count, err := o.QueryTable(user).Filter("name", s).All(&UsersShow); err == nil {
		beego.Info(count, UsersShow)
	} else {
		beego.Info("查询失败")
	}
	// err := o.Read(&user, "Name")
	// if err != nil {
	// 	beego.Info(err)
	// }
	// UsersShow = append(UsersShow, user)

	return UsersShow
}

//以姓名进行查询方法
func UsersRead4(s string) []WsUsers {
	var UsersShow []WsUsers
	user := WsUsers{}
	o := orm.NewOrm()
	if count, err := o.QueryTable(user).Filter("realname", s).All(&UsersShow); err == nil {
		beego.Info(count, UsersShow)
	} else {
		beego.Info("查询失败")
	}
	return UsersShow
}

//添加完用户 跳转至用户的编辑页面获取ID方法
func UsersRead5(s string) string {
	user := WsUsers{Name: s}
	o := orm.NewOrm()
	err := o.Read(&user, "Name")
	if err != nil {
		beego.Info(err)
	}
	id := strconv.Itoa(user.Id)
	url := "/admin/users/detail?user_id=" + id
	return url
}
