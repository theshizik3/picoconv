package configs

import (
	"converter/pkg/converter"
	"io"
)

type Decoder interface {
	Decode(r io.Reader, opts converter.Options) (any, error)
	Type() string
}
