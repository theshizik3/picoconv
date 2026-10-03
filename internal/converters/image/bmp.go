package image

import (
	"converter/pkg/converter"
	"image"
	"io"

	"golang.org/x/image/bmp"
)

type BMPDecoder struct{}

type BMPEncoder struct{}

func (d *BMPDecoder) Type() string {
	return "bmp"
}

func (e *BMPEncoder) Type() string {
	return "bmp"
}

func (d *BMPDecoder) Decode(r io.Reader) (image.Image, error) {
	img, err := bmp.Decode(r)
	return img, err
}

func (e *BMPEncoder) Encode(w io.Writer, img image.Image, opts converter.Options) error {
	err := bmp.Encode(w, img)
	return err
}
