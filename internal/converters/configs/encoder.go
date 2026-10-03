package configs

import (
	"io"

	"github.com/theshizik3/picoconv/pkg/converter"
)

type Encoder interface {
	Encode(w io.Writer, data any, opts converter.Options) error
	Type() string
}
