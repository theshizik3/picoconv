package configs

import (
	"converter/pkg/converter"
	"io"
)

type Encoder interface {
	Encode(w io.Writer, data any, opts converter.Options) error
	Type() string
}
