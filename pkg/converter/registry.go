package converter

type Registry interface {
	Register(c Converter) error
	Get(source, target string) (Converter, error)
	SupportedTypes() [][2]string
}
