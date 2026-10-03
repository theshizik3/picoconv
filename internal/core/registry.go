package core

import (
	"converter/pkg/converter"
	"fmt"
	"sync"
)

type key struct {
	src string
	trg string
}

type Registry struct {
	mu sync.RWMutex
	m  map[key]converter.Converter
}

func NewRegistry() *Registry {
	reg := Registry{m: make(map[key]converter.Converter)}
	return &reg
}

func (r *Registry) Register(conv converter.Converter) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	k := key{conv.Source(), conv.Target()}

	if _, exists := r.m[k]; exists {
		return fmt.Errorf("duplicate: %s -> %s", k.src, k.trg)
	}

	r.m[k] = conv
	return nil
}

func (r *Registry) Get(source, target string) (converter.Converter, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	conv, ok := r.m[key{src: source, trg: target}]
	if !ok {
		return nil, fmt.Errorf("converter %s -> %s does not exist", source, target)
	}

	return conv, nil
}

func (r *Registry) SupportedTypes() [][2]string {
	formats := [][2]string{}
	for k := range r.m {
		format := [2]string{}
		format[0] = k.src
		format[1] = k.trg
		formats = append(formats, format)
	}
	return formats
}
