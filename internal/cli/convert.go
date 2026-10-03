package cli

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/theshizik3/picoconv/internal/converters"
	"github.com/theshizik3/picoconv/internal/core"
	"github.com/theshizik3/picoconv/pkg/converter"

	"github.com/spf13/cobra"
)

var output string

var convertCmd = cobra.Command{
	Use:   "convert [flags] file",
	Short: "Converts the file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return convert(opts, src, trg, args[0], output)
	},
}

func init() {
	rootCmd.AddCommand(&convertCmd)

	convertCmd.Flags().StringVar(&output, "output", "", "output file")
	registerOptionsFlags(&convertCmd)
}

func convert(opts converter.Options, src, trg, input, output string) error {
	if trg == "" {
		return fmt.Errorf("output format not specified")
	}

	reg, err := converters.BuildRegistry()
	if err != nil {
		return err
	}

	if src == "" {
		src, err = core.GetFileType(input)
		if err != nil {
			return err
		}
	}

	if output == "" {
		output = strings.TrimSuffix(input, filepath.Ext(input)) + "." + trg
	}

	conv, err := reg.Get(src, trg)

	if err != nil {
		return err
	}

	err = conv.Convert(input, output, opts)

	return err
}
