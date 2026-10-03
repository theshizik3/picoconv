package cli

import (
	"github.com/theshizik3/picoconv/pkg/converter"

	"github.com/spf13/cobra"
)

var opts converter.Options
var src, trg string

func registerOptionsFlags(cmd *cobra.Command) {
	f := cmd.Flags()
	f.StringVar(&src, "from", "auto", "input file format")
	f.StringVar(&trg, "to", "", "output file format")
	f.IntVar(&opts.Quality, "quality", 80, "quality")
	f.IntVar(&opts.AlphaQuality, "alpha-quality", 80, "alpha channel quality")
	f.IntVar(&opts.Compress_lvl, "compress-lvl", 0, "compress level")
	f.IntVar(&opts.Method, "method", 4, "method")
	f.IntVar(&opts.Speed, "speed", 6, "speed")
	f.IntVar(&opts.Indent, "indent", 0, "indent")
	f.StringVar(&opts.Comma, "comma", ";", "csv comma")
	f.StringVar(&opts.Comment, "comment", "", "csv comment")
	f.BoolVar(&opts.Lossless, "lossless", false, "lossless")
}
