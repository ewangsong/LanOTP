package radius

import (
	"database/sql"
	"fmt"

	"github.com/astaxie/beego"

	_ "github.com/Go-SQL-Driver/MySQL"
)

var a, b = GetDbConfig()

//判断是否正确连接数据库
func DbLive() {
	fmt.Println(a, b)
	db, err := sql.Open(a, b)
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
	db, _ := sql.Open(a, b)
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

//从数据库中读取用户名和密码
func GetUserPasswd(u string) (p string) {
	db, _ := sql.Open(a, b)
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
