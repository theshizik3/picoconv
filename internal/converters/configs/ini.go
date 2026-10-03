package configs

import (
	"converter/pkg/converter"
	"io"

	"github.com/ltick/go-ini"
)

type INIDecoder struct{}

type INIEncoder struct{}

func (d *INIDecoder) Type() string {
	return "ini"
}

func (e *INIEncoder) Type() string {
	return "ini"
}

func (d *INIDecoder) Decode(r io.Reader, opts converter.Options) (any, error) {
	b := []byte{}
	_, err := r.Read(b)
	if err != nil {
		return nil, err
	}
	var data any
	err = ini.Unmarshal(b, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (e *INIEncoder) Encode(w io.Writer, data any, opts converter.Options) error {
	b, err := ini.Marshal(data)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}
