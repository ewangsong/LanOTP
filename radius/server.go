package radius

import (
	"net"

	"github.com/astaxie/beego"
)

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
			if GetUserPasswd(userName) == uerPassword {
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
