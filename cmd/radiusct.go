package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"radiusweb/radius"

	"github.com/spf13/cobra"
)

func init() {
	var daemon bool
	startCmd := &cobra.Command{
		Use:   "radiusct",
		Short: "start lanradius",

		Run: func(cmd *cobra.Command, args []string) {
			if daemon {
				command := exec.Command("/opt/lanradius/lanradius", "radiusct")
				command.Start()
				fmt.Printf("lanradius radiusct start, [PID] %d running...\n", command.Process.Pid)
				ioutil.WriteFile("/run/lanradius-radiusct.pid", []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)
				daemon = false
				os.Exit(0)
			} else {
				fmt.Println("lanradius radiusct start")
			}
			radius.RadiusRun()
		},
	}
	startCmd.Flags().BoolVarP(&daemon, "deamon", "d", false, "is daemon?")
	RootCmd.AddCommand(startCmd)

}
