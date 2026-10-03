package converters

import (
	"converter/internal/converters/audio"
	"converter/internal/converters/configs"
	"converter/internal/converters/image"
	"converter/internal/converters/video"
	"converter/internal/core"
	"converter/pkg/converter"
)

func BuildRegistry() (converter.Registry, error) {
	reg := core.NewRegistry()

	{ // image converters
		image_encoders := map[string]image.Encoder{
			"jpeg": &image.JPEGEncoder{},
			"png":  &image.PNGEncoder{},
			"webp": &image.WEBPEncoder{},
			"avif": &image.AVIFEncoder{},
			"bmp":  &image.BMPEncoder{},
			"tiff": &image.TIFFEncoder{},
		}
		image_decoders := map[string]image.Decoder{
			"jpeg": &image.JPEGDecoder{},
			"png":  &image.PNGDecoder{},
			"webp": &image.WEBPDecoder{},
			"avif": &image.AVIFDecoder{},
			"bmp":  &image.BMPDecoder{},
			"heic": &image.HEICDecoder{},
			"tiff": &image.TIFFDecoder{},
			"svg":  &image.SVGDecoder{},
		}
		for e := range image_encoders {
			for d := range image_decoders {
				if e != d {
					err := reg.Register(&image.ImageConverter{Enc: image_encoders[e], Dec: image_decoders[d]})
					if err != nil {
						return nil, err
					}
				}
			}
		}
	}

	{ // video formats
		video_formats := []string{
			"mp4",
			"mkv",
			"mov",
			"avi",
			"webm",
			"wmv",
			"flv",
		}
		for _, src := range video_formats {
			for _, trg := range video_formats {
				if src != trg {
					err := reg.Register(&video.VideoConverter{Src: src, Trg: trg})
					if err != nil {
						return nil, err
					}
				}
			}
		}
	}

	{ // audio formats
		audio_formats := []string{
			"mp3",
			"flac",
			"wav",
			"ogg",
			"aac",
			"opus",
		}
		for _, src := range audio_formats {
			for _, trg := range audio_formats {
				if src != trg {
					err := reg.Register(&audio.AudioConverter{Src: src, Trg: trg})
					if err != nil {
						return nil, err
					}
				}
			}
		}
	}

	{ // configs formats
		config_encoders := []configs.Encoder{
			&configs.INIEncoder{},
			&configs.JSONEncoder{},
			&configs.TOMLEncoder{},
			&configs.YAMLEncoder{},
		}
		config_decoders := []configs.Decoder{
			&configs.INIDecoder{},
			&configs.JSONDecoder{},
			&configs.TOMLDecoder{},
			&configs.YAMLDecoder{},
		}

		for e := range config_encoders {
			for d := range config_decoders {
				if e != d {
					err := reg.Register(&configs.ConfigsConverter{Enc: config_encoders[e], Dec: config_decoders[d]})
					if err != nil {
						return nil, err
					}
				}
			}
		}
	}

	return reg, nil
}
