package image

import (
	"fmt"
	"image"
	"image/png"
	"io"

	"github.com/theshizik3/picoconv/pkg/converter"
)

type PNGDecoder struct{}

type PNGEncoder struct{}

func (d *PNGDecoder) Type() string {
	return "png"
}

func (e *PNGEncoder) Type() string {
	return "png"
}

func (d *PNGDecoder) Decode(r io.Reader) (image.Image, error) {
	img, err := png.Decode(r)
	return img, err
}

func (e *PNGEncoder) Encode(w io.Writer, img image.Image, opts converter.Options) error {
	compress_lvl := opts.Compress_lvl

	if (compress_lvl < -3) || (compress_lvl > 0) {
		return fmt.Errorf("invalid compression level")
	}

	encoder := png.Encoder{CompressionLevel: png.CompressionLevel(compress_lvl)}
	err := encoder.Encode(w, img)

	return err
}
