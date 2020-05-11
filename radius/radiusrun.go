package radius

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/astaxie/beego"
)

func RadiusRun() {
	logfile, err := os.OpenFile("/var/log/lanradius/goradius.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		os.Exit(-1)
	}
	defer logfile.Close()
	log.SetOutput(logfile)
	log.SetPrefix("[INFO]")
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	//	goradius.DbLive() //判断数据库是否存在
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
