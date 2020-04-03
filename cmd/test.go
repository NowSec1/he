package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")
		g1 := make([]byte, 1024*1024*1024*10)
		fmt.Println("cap(1G):", cap(g1), "len(1G):", len(g1))
		for i := 0; i < cap(g1); i++ {
			g1[i] = byte(i)
		}
		time.Sleep(time.Minute)
		for i := range g1 {
			fmt.Println(i)
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
