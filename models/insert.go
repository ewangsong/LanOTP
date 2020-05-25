package models

import (
	"ewangsong/LanOTP/otp"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

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

//BasInsert 添加
func BasInsert(name, ip_add, secret, port string) (bas WsBas, err error) {
	bas = WsBas{Name: name,
		IpAddr: ip_add,
		Secret: secret,
		Port:   port}
	o := orm.NewOrm()
	_, err = o.Insert(&bas)
	if err != nil {
		beego.Info("bas插入错误", err)
		return bas, err
	} else {
		return bas, err
	}

}

//添加用户数据库操作
func UserInsert(realname, name string) (user WsUsers, err error) {
	user = WsUsers{RealName: realname, Name: name}
	o := orm.NewOrm()
	if realname == "" || name == "" {
		err = fmt.Errorf("用户名或者姓名为空")
		return user, err
	}
	err = o.Read(&user, "Name") //先查询是否存在此用户
	if err != nil {             //没有查询到
		_, err1 := o.Insert(&user)
		if err1 != nil {
			beego.Info("添加用户插入错误", err1)
			return
		}
	} else { //账号已存在
		return user, fmt.Errorf("账号已存在")
	}
	return user, nil
}

//token 添加
func AddToken(username string, typeid int) (token WsOtp, b bool) {
	token = WsOtp{
		Secret:      otp.GenerateRandomSecret(),
		OtpType:     typeid,
		BindingUser: username,
	}
	if username == "" {
		return token, false
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
		return token, false
	}
	err = o.Read(&token, "Secret")
	if err != nil {
		beego.Info(err)
		return token, false
	}
	token.OtpSn = "LANOTP" + strconv.Itoa(token.Id)
	_, err = o.Update(&token)
	if err != nil {
		beego.Info(err)
		return token, false
	}
	return token, true
}

//添加操作日志
func LogInsert(name, ip, desc string) bool {
	wslog := WsLog{
		OperatorNmae: name,
		OperatIp:     ip,
		OperatDesc:   desc,
	}
	o := orm.NewOrm()
	_, err := o.Insert(&wslog)

	if err != nil {
		beego.Error("插入日历错误err：", err)
		return false
	}
	return true
}
