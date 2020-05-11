package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "lanradius",
	Short: "Lanradius认证系统",
	Long:  "Lanradius主要是为懒投资及所属公司的一套认证系统，包括MAC地址认证，OTP二次认证",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`admin			Start Lanradius`)
		fmt.Println(`help        	Help about any command`)
		fmt.Println(`radiusct    	Start Lanradius`)
		fmt.Println(`stop        	Stop Lanradius`)
		fmt.Println(`version     	Print the version number of Lanradius`)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
