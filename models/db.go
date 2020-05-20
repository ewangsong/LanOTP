package models

import (
	"time"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//管理用户
type WsAdmin struct {
	Id       int `orm:"pk;auto"`
	Name     string
	Password string
}

//otp信息
type WsOtp struct {
	Id          int `orm:"auto;pk"`
	OtpSn       string
	OtpType     int
	Secret      string
	BindingUser string    `orm:"null"`
	Counter     uint64    `orm:"null"`
	OperatTime  time.Time `orm:"auto_now;type(datetime)"`
}

//用户信息
type WsUsers struct {
	Id           int `orm:"pk;auto"`
	Name         string
	RealName     string
	BindingToken int       `orm:"null"`
	CreateTime   time.Time `orm:"auto_now;type(datetime)"`
}

//bas信息
type WsBas struct {
	Id     int `orm:"pk;auto"`
	Name   string
	IpAddr string
	Secret string
	Port   string
}

//日志信息
type WsLog struct {
	Id           int `orm:"pk;auto"`
	OperatorNmae string
	OperatIp     string
	OperatTime   time.Time `orm:"auto_now;type(datetime)"`
	OperatDesc   string
}

func init() {
	dbtype := beego.AppConfig.String("dbtype")
	dbinfo := beego.AppConfig.String("dbinfo")
	orm.RegisterDataBase("default", dbtype, dbinfo)
	orm.RegisterModel(new(WsOtp), new(WsUsers), new(WsAdmin), new(WsBas), new(WsLog))
	orm.RunSyncdb("default", false, false)
	AddAdmin() //添加默认管理员账号
}
