package converter

type Converter interface {
	Source() string
	Target() string
	Convert(input, output string, opts Options) error
}
