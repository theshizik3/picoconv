package image

import (
	"fmt"
	"image"
	"io"

	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

type SVGDecoder struct{}

func (d *SVGDecoder) Type() string {
	return "svg"
}

func (d *SVGDecoder) Decode(r io.Reader) (image.Image, error) {
	icon, err := oksvg.ReadIconStream(r)
	if err != nil {
		return nil, err
	}

	w := int(icon.ViewBox.W)
	h := int(icon.ViewBox.H)

	if (w == 0) || (h == 0) {
		return nil, fmt.Errorf("viewbox is not defined")
	}

	icon.SetTarget(0, 0, float64(w), float64(h))

	img := image.NewRGBA(image.Rect(0, 0, w, h))
	dasher := rasterx.NewDasher(w, h, rasterx.NewScannerGV(w, h, img, img.Bounds()))
	icon.Draw(dasher, 1.0)

	return img, nil
}
