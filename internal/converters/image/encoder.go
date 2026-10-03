package image

import (
	"image"
	"io"

	"github.com/theshizik3/picoconv/pkg/converter"
)

type Encoder interface {
	Encode(w io.Writer, img image.Image, opts converter.Options) error
	Type() string
}
