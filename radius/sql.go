package radius

import (
	"database/sql"
	"ewangsong/LanOTP/otp"

	"github.com/astaxie/beego"

	_ "github.com/Go-SQL-Driver/MySQL"
)

var dbtype, dbinfo = GetDbConfig()

//判断是否正确连接数据库
func DbLive() {
	db, err := sql.Open(dbtype, dbinfo)
	defer db.Close()
	if err != nil {
		beego.Error(err)
	}
	// err1 := db.Ping()
	// if err1 != nil {
	// 	log.Fatal("Wrong connect to DB")

	// }
}

//由客户端IP得到客户端对应得密钥
func GetSecretAndDiff(ipp string) (secret []byte) {
	db, _ := sql.Open(dbtype, dbinfo)
	defer db.Close()
	res := db.QueryRow("select ip_addr,secret from ws_bas where ip_addr = ?", ipp)
	err2 := res.Scan(&ipp, &secret)
	if err2 != nil {
		beego.Error(err2)
	}
	if string(secret) == "" || string(ipp) == "" {
		beego.Info("请把IP和密钥加入认证服务器白名单")
		return nil
	} else {
		secret = []byte(secret)
		return secret
	}

}

//由用户名从数据库中读取用户名和密码
func GetUserPasswd(u string) (p string) {
	db, _ := sql.Open(dbtype, dbinfo)
	defer db.Close()
	var password string
	err2 := db.QueryRow("select password from ws_users where name = ?", u).Scan(&password)
	if err2 != nil {
		beego.Error(err2)
		return ""
	} else {
		return password
	}

}

//GetUserToken 由用户名得到otp token
func GetUserToken(user string) (p string) {
	db, err := sql.Open(dbtype, dbinfo)
	if err != nil {
		beego.Error("打开数据库错误err：")
	}
	defer db.Close()
	var totp otp.TOTP
	err = db.QueryRow("select secret from ws_otp where binding_user=?", user).Scan(&totp.Secret)
	if err != nil {
		beego.Error("查询用户错误", user)
		return
	}
	return totp.TotpGet()

}
