package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cryptoCmd represents the crypto command
var cryptoCmd = &cobra.Command{
	Use:   "crypto",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("crypto called")
	},
}

var cryptoMD5Cmd = &cobra.Command{
	Use:   "md5",
	Short: "md5",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("crypto md5 called")

	},
}

func init() {
	rootCmd.AddCommand(cryptoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cryptoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cryptoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
