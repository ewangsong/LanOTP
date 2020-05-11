package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/astaxie/beego"
	"github.com/spf13/cobra"
)

func init() {
	var daemon bool
	startCmd := &cobra.Command{
		Use:   "admin",
		Short: "start lanradius",

		Run: func(cmd *cobra.Command, args []string) {
			if daemon {
				command := exec.Command("/opt/lanradius/lanradius", "admin")
				command.Start()
				fmt.Printf("lanradius admin start, [PID] %d running...\n", command.Process.Pid)
				ioutil.WriteFile("/run/lanradius-admin.pid", []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)
				daemon = false
				os.Exit(0)
			} else {
				fmt.Println("lanradius admin start")
			}
			beego.Run()
		},
	}
	startCmd.Flags().BoolVarP(&daemon, "deamon", "d", false, "is daemon?")
	RootCmd.AddCommand(startCmd)

}
