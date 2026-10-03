package image

import (
	"image"
	"io"

	"github.com/adrium/goheif"
)

type HEICDecoder struct{}

func (d *HEICDecoder) Type() string {
	return "heic"
}

func (d *HEICDecoder) Decode(r io.Reader) (image.Image, error) {
	img, err := goheif.Decode(r)
	return img, err
}
