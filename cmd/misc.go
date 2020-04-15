package cmd

import (
	"fmt"

	"github.com/zu1k/he/modules/png"

	"github.com/spf13/cobra"
)

// miscCmd represents the misc command
var miscCmd = &cobra.Command{
	Use:   "misc",
	Short: "misc tools for ctf",
	Long: `Some Misc tools for ctf.

Now include:
	PNG crc check.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("misc called")
	},
}

var PngCRCCmd = &cobra.Command{
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
			fmt.Println("Help: with -h flag for help")
		}
	},
}

func init() {
	rootCmd.AddCommand(miscCmd)
	miscCmd.AddCommand(PngCRCCmd)
	PngCRCCmd.Flags().StringP("file", "f", "", "png file path")
	PngCRCCmd.Flags().BoolP("fix", "x", false, "fix png file")
}
