package cmd

import (
	"fmt"

	"github.com/khanghldk/gokit/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// verCmd represents the version command
var verCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Get version",
	Run: func(cmd *cobra.Command, args []string) {
		ver, err  := utils.GetVersion()
		if err != nil {
			logrus.Error("invalid version")
			return
		}

		fmt.Printf("gokit version %v", ver)
	},
}

func init() {
	RootCmd.AddCommand(verCmd)
}

