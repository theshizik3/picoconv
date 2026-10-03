package configs

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/theshizik3/picoconv/pkg/converter"
)

type JSONDecoder struct{}

type JSONEncoder struct{}

func (d *JSONDecoder) Type() string {
	return "json"
}

func (e *JSONEncoder) Type() string {
	return "json"
}

func (d *JSONDecoder) Decode(r io.Reader, opts converter.Options) (any, error) {
	b := []byte{}
	_, err := r.Read(b)
	if err != nil {
		return nil, err
	}
	var data any
	err = json.Unmarshal(b, &data)

	if err != nil {
		return nil, err
	}
	return data, nil
}

func (e *JSONEncoder) Encode(w io.Writer, data any, opts converter.Options) error {
	indent := opts.Indent
	b, err := json.MarshalIndent(data, "", strings.Repeat(" ", indent))
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}
