package configs

import (
	"io"

	"github.com/theshizik3/picoconv/pkg/converter"

	"gopkg.in/yaml.v3"
)

type YAMLDecoder struct{}

type YAMLEncoder struct{}

func (d *YAMLDecoder) Type() string {
	return "yaml"
}

func (e *YAMLEncoder) Type() string {
	return "yaml"
}

func (d *YAMLDecoder) Decode(r io.Reader, opts converter.Options) (any, error) {
	b := []byte{}
	_, err := r.Read(b)
	if err != nil {
		return nil, err
	}
	var data any
	err = yaml.Unmarshal(b, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (e *YAMLEncoder) Encode(w io.Writer, data any, opts converter.Options) error {
	b, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}
