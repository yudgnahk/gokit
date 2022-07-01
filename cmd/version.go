package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yudgnahk/gokit/constants"
)

// verCmd represents the version command
var verCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Get version",
	Run: func(cmd *cobra.Command, args []string) {
		message := fmt.Sprintf("gokit version %v", constants.Version)
		fmt.Println(constants.ColorYellow, message)
	},
}

func init() {
	RootCmd.AddCommand(verCmd)
}
