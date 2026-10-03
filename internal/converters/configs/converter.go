package configs

import (
	"os"

	"github.com/theshizik3/picoconv/internal/core"
	"github.com/theshizik3/picoconv/pkg/converter"
)

type ConfigsConverter struct {
	core.BaseConverter
	Dec Decoder
	Enc Encoder
}

func (c *ConfigsConverter) Source() string {
	return c.Dec.Type()
}

func (c *ConfigsConverter) Target() string {
	return c.Enc.Type()
}

func (c *ConfigsConverter) Convert(input, output string, opts converter.Options) error {
	in_file, err := os.OpenFile(input, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	m, err := c.Dec.Decode(in_file, opts)
	if err != nil {
		return err
	}
	out_file, err := os.Create(output)
	if err != nil {
		return err
	}
	err = c.Enc.Encode(out_file, m, opts)
	return err
}
