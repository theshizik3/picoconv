package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = cobra.Command{
	Use:   "version",
	Short: "Displays the installed version of PicoConv",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("PicoConv version 0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(&versionCmd)
}
