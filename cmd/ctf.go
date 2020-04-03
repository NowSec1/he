package cmd

import (
	"fmt"

	"github.com/zu1k/he/modules/png"

	"github.com/spf13/cobra"
)

// ctfCmd represents the ctf command
var ctfCmd = &cobra.Command{
	Use:   "ctf",
	Short: "some tools for ctf",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ctf called")
	},
}

var ctfMiscCmd = &cobra.Command{
	Use:   "misc",
	Short: "some misc tools for ctf",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ctf misc called")
	},
}

var ctfMiscPngCRCCmd = &cobra.Command{
	Use:   "pngcrc",
	Short: "check and fix png crc",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if f := cmd.Flag("file"); f.Changed {
			filePath := f.Value.String()
			fmt.Println("Checking PNG file...")
			pngfile, err := png.NewPNG(filePath)
			if err != nil {
				fmt.Println(err.Error())
			}
			fix := false
			fix, _ = cmd.Flags().GetBool("fix")
			pngfile.CheckAllChunks(fix)
		} else {
			fmt.Println("Help: he ctf misc pngcrc -h")
		}
	},
}

func init() {
	rootCmd.AddCommand(ctfCmd)
	ctfCmd.AddCommand(ctfMiscCmd)
	ctfMiscCmd.AddCommand(ctfMiscPngCRCCmd)
	ctfMiscPngCRCCmd.Flags().StringP("file", "f", "", "png file path")
	ctfMiscPngCRCCmd.Flags().BoolP("fix", "x", false, "fix png file")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ctfCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ctfCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
