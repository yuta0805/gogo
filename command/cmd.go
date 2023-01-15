package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "tfmake",
	Short: "command line is terraform file maker",
	Run: func (cmd *cobra.Command, arg []string)  {
		fmt.Println("main command")
	},
}


func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(
		makeCommand(),
	)
}
