package models

import (
	"time"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type WsUsers struct {
	Id         int `orm:"pk;auto"`
	TypeId     int `orm:"null"`
	Name       string
	Password   string
	RealName   string
	CreateTime time.Time `orm:"auto_now;type(datetime)"`
}
type WsAdmin struct {
	Id       int `orm:"pk;auto"`
	Name     string
	Password string
}
type WsBas struct {
	Id     int `orm:"pk;auto"`
	Name   string
	IpAddr string
	Secret string
	Port   string
}

type WsLog struct {
	Id           int `orm:"pk;auto"`
	OperatorNmae string
	OperatIp     string
	OperatTime   time.Time `orm:"auto_now;type(datetime)"`
	OperatDesc   string
}

type WsOtp struct {
	OtpId       int `orm:"pk"`
	Secret      string
	BindingUser string
	OperatTime  time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	// user := beego.AppConfig.String("mysqluser")
	// password := beego.AppConfig.String("mysqlpass")
	// ip := beego.AppConfig.String("mysqlurls")
	// mydb := beego.AppConfig.String("mysqldb")
	// dbconfig := user + ":" + password + "@tcp" + "(" + ip + ":" + "3306" + ")" + "/" + mydb + "?" + "charset=utf8&loc=Local"
	dbtype := beego.AppConfig.String("dbtype")
	dbinfo := beego.AppConfig.String("dbinfo")

	orm.RegisterDataBase("default", dbtype, dbinfo)
	orm.RegisterModel(new(WsUsers), new(WsAdmin), new(WsBas), new(WsLog), new(WsOtp))
	orm.RunSyncdb("default", false, false)
	AddAdmin() //添加默认管理员账号
}
