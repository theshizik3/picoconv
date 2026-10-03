package image

import (
	"image"
	"io"

	"github.com/theshizik3/picoconv/pkg/converter"

	"fmt"

	"github.com/vegidio/avif-go"
)

type AVIFDecoder struct{}

type AVIFEncoder struct{}

func (d *AVIFDecoder) Type() string {
	return "avif"
}

func (e *AVIFEncoder) Type() string {
	return "avif"
}

func (d *AVIFDecoder) Decode(r io.Reader) (image.Image, error) {
	img, err := avif.Decode(r)
	return img, err
}

func (e *AVIFEncoder) Encode(w io.Writer, img image.Image, opts converter.Options) error {
	quality := opts.Quality

	if (quality > 100) || (quality < 0) {
		return fmt.Errorf("invalid quality")
	}

	alpha_quality := opts.AlphaQuality

	if (alpha_quality > 100) || (alpha_quality < 0) {
		return fmt.Errorf("invalid alpha quality")
	}

	speed := opts.Speed

	if (speed < 0) || (speed > 10) {
		return fmt.Errorf("invalid speed")
	}

	err := avif.Encode(w, img, &avif.Options{
		Speed:        speed,
		ColorQuality: quality,
		AlphaQuality: alpha_quality,
	})

	return err
}
