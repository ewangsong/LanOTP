package cmd

import (
	"ewangsong/LanOTP/radius"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

var (
	start  bool
	stop   bool
	v      bool
	h      bool
	startd bool
)

//Cmd 命令行
func Cmd() {
	flag.BoolVar(&start, "start", false, "前台启动")
	flag.BoolVar(&stop, "stop", false, "关闭程序")
	flag.BoolVar(&v, "v", false, "查看版本")
	flag.BoolVar(&startd, "startd", false, "后台启动")

	flag.Parse()
	command := exec.Command("./lanotp", "-start")

	if start {
		radius.RadiusRun()
	}

	if stop {
		fmt.Println("关闭程序")
		stopp()
	}
	if v {
		fmt.Println("当前版本是0.5.0")
	}

	if startd {
		fmt.Println("后台启动")
		err := command.Start()
		fmt.Printf("gonne start, [PID] %d running...\n", command.Process.Pid)
		ioutil.WriteFile("/run/lanotp.pid", []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)
		if err != nil {
			fmt.Println("启动程序失败", err)
			return
		}

	} else {
		command.Wait()

	}
	//打印默认帮助
	if len(os.Args) == 1 {
		flag.PrintDefaults()
	}

}

func stopp() {
	b, err := ioutil.ReadFile("/run/lanotp.pid")

	if err != nil {
		fmt.Println("获取程序PID错误", err)
		return
	}

	stopcommand := exec.Command("/bin/kill", string(b))

	stopcommand.Start()

}
