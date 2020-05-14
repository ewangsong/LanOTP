package models

import (
	"ewangsong/LanOTP/otp"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//BasInsert 添加
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

//添加用户数据库操作

func UserInsert(realname, name string) bool {
	user := WsUsers{RealName: realname, Name: name}
	o := orm.NewOrm()
	if realname == "" || name == "" {
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

//token 添加
func AddToken(username string, typeid int) (tokenid int, b bool) {
	token := WsOtp{
		Secret:      otp.GenerateRandomSecret(),
		OtpType:     typeid,
		BindingUser: username,
	}
	if username == "" {
		return 0, false
	}
	o := orm.NewOrm()
	err := o.Read(&token, "Secret")
	for {
		if err != nil { //不存在
			break
		} else { //存在
			token.Secret = otp.GenerateRandomSecret()
			err = o.Read(&token, "Secret")
		}
	}

	_, err = o.Insert(&token)
	if err != nil {
		beego.Info("用户插入错误", err)
		return 0, false
	}
	err = o.Read(&token, "Secret")
	if err != nil {
		beego.Info(err)
		return -1, false
	}
	token.OtpSn = "LANOTP" + strconv.Itoa(token.Id)
	_, err = o.Update(&token)
	if err != nil {
		beego.Info(err)
		return -1, false
	}

	return token.Id, true
}

// //插入操作日志
// func LogInsert() {

// }
