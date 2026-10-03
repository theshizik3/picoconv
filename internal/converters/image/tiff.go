package image

import (
	"converter/pkg/converter"
	"fmt"
	"image"
	"io"

	"github.com/sunshineplan/tiff"
)

type TIFFDecoder struct{}

type TIFFEncoder struct{}

func (d *TIFFDecoder) Type() string {
	return "tiff"
}

func (e *TIFFEncoder) Type() string {
	return "tiff"
}

func (d *TIFFDecoder) Decode(r io.Reader) (image.Image, error) {
	img, err := tiff.Decode(r)
	return img, err
}

func (e *TIFFEncoder) Encode(w io.Writer, img image.Image, opts converter.Options) error {
	compress_lvl := opts.Compress_lvl

	if (compress_lvl > 5) || (compress_lvl < 0) {
		return fmt.Errorf("invalid compression level")
	}

	err := tiff.Encode(w, img, &tiff.Options{
		Compression: tiff.CompressionType(compress_lvl),
	})

	return err
}
