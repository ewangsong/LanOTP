package radius

import "github.com/astaxie/beego"

//获取数据库配置
func GetDbConfig() (t, m string) {
	dbtype := beego.AppConfig.String("dbtype")
	dbinfo := beego.AppConfig.String("dbinfo")
	return dbtype, dbinfo
}
