package radius

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/astaxie/beego"
)

//RadiusRun 启动服务
func RadiusRun() {
	logfile, err := os.OpenFile("./goradius.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		os.Exit(-1)
	}
	defer logfile.Close()
	log.SetOutput(logfile)
	log.SetPrefix("[INFO]")
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
	port := beego.AppConfig.String("radiusport")
	udpaddr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		fmt.Println(err)
	}

	udpconn, err2 := net.ListenUDP("udp", udpaddr)
	if err2 != nil {
		beego.Error(err2)
	}

	RadiusServer(udpconn)
}

//RadiusServe 服务
func RadiusServer(conn *net.UDPConn) {

	for {
		buf := make([]byte, 4096)
		_, udpaddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			beego.Error(err)
		}

		ipp := udpaddr.IP.String()
		secret := GetSecretAndDiff(ipp)
		if secret == nil {
			beego.Info("密钥为空")
			continue
		}
		pakage, _ := Parse(buf, secret)
		userName := UserName_GetString(pakage)
		uerPassword := UserPassword_GetString(pakage)
		defer conn.Close()
		go func() {
			if GetUserToken(userName) == uerPassword {
				res := pakage.Response(CodeAccessAccept)
				var vl = []byte{'o', 'k'}
				ReplyMessage_Add(res, vl)
				udpp, _ := res.Encode()
				conn.WriteToUDP(udpp, udpaddr)
				//打印用户名密码和NAS IP
				beego.Info("username=", userName, "userpassword=", uerPassword, "nasIP=", ipp, "OK")

			} else {
				res := pakage.Response(CodeAccessReject)
				var vl = []byte{'n', 'o'}
				ReplyMessage_Add(res, vl)
				udpp, _ := res.Encode()
				conn.WriteToUDP(udpp, udpaddr)
				//打印用户名密码和NAS IP
				beego.Info("username=", userName, "userpassword=", uerPassword, "nasIP=", ipp, "NO")
			}

		}()
	}
}
