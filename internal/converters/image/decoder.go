package image

import (
	"image"
	"io"
)

type Decoder interface {
	Decode(r io.Reader) (image.Image, error)
	Type() string
}
