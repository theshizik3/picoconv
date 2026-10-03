package image

import (
	"converter/pkg/converter"
	"image"
	"io"
)

type Encoder interface {
	Encode(w io.Writer, img image.Image, opts converter.Options) error
	Type() string
}
