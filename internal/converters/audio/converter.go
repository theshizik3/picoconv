package audio

import (
	"github.com/theshizik3/picoconv/internal/core"
	"github.com/theshizik3/picoconv/pkg/converter"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type AudioConverter struct {
	core.BaseConverter
	Src string
	Trg string
}

func (c *AudioConverter) Source() string {
	return c.Src
}

func (c *AudioConverter) Target() string {
	return c.Trg
}

func (c *AudioConverter) Convert(input, output string, opts converter.Options) error {
	kwargs := ffmpeg.KwArgs{}
	kwargs["format"] = c.Target()

	err := ffmpeg.Input(input).
		Output(output, kwargs).
		GlobalArgs("-loglevel", "quiet").
		OverWriteOutput().
		Run()
	return err
}
