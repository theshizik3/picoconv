package configs

import (
	"io"

	"github.com/theshizik3/picoconv/pkg/converter"
)

type Decoder interface {
	Decode(r io.Reader, opts converter.Options) (any, error)
	Type() string
}
