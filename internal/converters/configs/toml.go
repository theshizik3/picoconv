package configs

import (
	"io"

	"github.com/theshizik3/picoconv/pkg/converter"

	"github.com/BurntSushi/toml"
)

type TOMLDecoder struct{}

type TOMLEncoder struct{}

func (d *TOMLDecoder) Type() string {
	return "toml"
}

func (e *TOMLEncoder) Type() string {
	return "toml"
}

func (d *TOMLDecoder) Decode(r io.Reader, opts converter.Options) (any, error) {
	b := []byte{}
	_, err := r.Read(b)
	if err != nil {
		return nil, err
	}
	var data any
	err = toml.Unmarshal(b, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (e *TOMLEncoder) Encode(w io.Writer, data any, opts converter.Options) error {
	b, err := toml.Marshal(data)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}
