package cli

import (
	"converter/internal/core"
	"fmt"

	"github.com/spf13/cobra"
)

var inspectCmd = cobra.Command{
	Use:   "inspect",
	Short: "Determines the file type",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := core.GetFileType(args[0])
		if err != nil {
			return err
		}
		fmt.Println(f)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(&inspectCmd)
}
