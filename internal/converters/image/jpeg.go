package image

import (
	"converter/pkg/converter"
	"fmt"
	"image"
	"image/jpeg"
	"io"
)

type JPEGDecoder struct{}

type JPEGEncoder struct{}

func (d *JPEGDecoder) Type() string {
	return "jpeg"
}

func (e *JPEGEncoder) Type() string {
	return "jpeg"
}

func (d *JPEGDecoder) Decode(r io.Reader) (image.Image, error) {
	img, _, err := image.Decode(r)
	return img, err
}

func (e *JPEGEncoder) Encode(w io.Writer, img image.Image, opts converter.Options) error {
	quality := opts.Quality

	if (quality > 100) || (quality < 0) {
		return fmt.Errorf("invalid quality level")
	}

	options := &jpeg.Options{Quality: quality}

	err := jpeg.Encode(w, img, options)

	return err
}
