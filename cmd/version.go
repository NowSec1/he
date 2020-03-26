package cmd

import (
	"fmt"

	"github.com/zu1k/he/constant"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show He version",
	Long:  `Show the He tools version information.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(constant.Name, "-", constant.Version, "-", constant.BuildTime, "\n")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
