package cli

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pconv",
	Short: "Lightweight media converter",
	Long:  `PicoConv is an open-source media converter capable of handling various video, audio, and image formats`,
}

func Execute() {
	log.SetOutput(io.Discard)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
