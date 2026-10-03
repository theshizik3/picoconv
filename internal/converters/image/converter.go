package image

import (
	"os"

	"github.com/theshizik3/picoconv/internal/core"
	"github.com/theshizik3/picoconv/pkg/converter"
)

type ImageConverter struct {
	core.BaseConverter
	Dec Decoder
	Enc Encoder
}

func (c *ImageConverter) Source() string {
	return c.Dec.Type()
}

func (c *ImageConverter) Target() string {
	return c.Enc.Type()
}

func (c *ImageConverter) Convert(input, output string, opts converter.Options) error {
	in_file, err := os.OpenFile(input, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	img, err := c.Dec.Decode(in_file)
	if err != nil {
		return err
	}
	out_file, err := os.Create(output)
	if err != nil {
		return err
	}
	err = c.Enc.Encode(out_file, img, opts)
	return err
}
