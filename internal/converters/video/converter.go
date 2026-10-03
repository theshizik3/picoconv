package video

import (
	"github.com/theshizik3/picoconv/internal/core"
	"github.com/theshizik3/picoconv/pkg/converter"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type VideoConverter struct {
	core.BaseConverter
	Src string
	Trg string
}

func (c *VideoConverter) Source() string {
	return c.Src
}

func (c *VideoConverter) Target() string {
	return c.Trg
}

func (c *VideoConverter) Convert(input, output string, opts converter.Options) error {
	kwargs := ffmpeg.KwArgs{}
	if c.Target() == "gif" {
		kwargs["vf"] = "split[s0][s1];[s0]palettegen[p];[s1][p]paletteuse"
	}
	kwargs["format"] = c.Target()

	err := ffmpeg.Input(input).
		Output(output, kwargs).
		GlobalArgs("-loglevel", "quiet").
		OverWriteOutput().
		Run()
	return err
}
