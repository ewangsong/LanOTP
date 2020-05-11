package cmd

import "github.com/spf13/cobra"

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Lanradius",
	Run: func(cmd *cobra.Command, args []string) {
		println("Lanradius version is 0.1.0")
	},
}
