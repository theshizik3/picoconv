package image

import (
	"fmt"
	"image"
	"io"

	"github.com/theshizik3/picoconv/pkg/converter"

	"github.com/deepteams/webp"
)

type WEBPDecoder struct{}

type WEBPEncoder struct{}

func (d *WEBPDecoder) Type() string {
	return "webp"
}

func (e *WEBPEncoder) Type() string {
	return "webp"
}

func (d *WEBPDecoder) Decode(r io.Reader) (image.Image, error) {
	img, err := webp.Decode(r)
	return img, err
}

func (e *WEBPEncoder) Encode(w io.Writer, img image.Image, opts converter.Options) error {
	lossless := opts.Lossless
	quality := opts.Quality

	if (quality > 100) || (quality < 0) {
		return fmt.Errorf("invalid quality")
	}

	method := opts.Method

	if (method < 0) || (method > 6) {
		return fmt.Errorf("invalid method")
	}

	err := webp.Encode(w, img, &webp.EncoderOptions{
		Lossless: lossless,
		Quality:  float32(quality),
		Method:   method,
	})

	return err
}
