package mouse

import (
	"errors"
)

type Mouse_memory struct {
	data map[string][]byte
}

func (m *Mouse_memory) Put(key string, value []byte) error {
	m.data[key] = value
	return nil
}

func (m *Mouse_memory) Get(key string) ([]byte, error) {
	value, ok := m.data[key]
	if !ok {
		return nil, errors.New("key not found")
	}
	return value, nil
}

func Mouse_memory_new() *Mouse_memory {
	return &Mouse_memory{
		data: map[string][]byte{},
	}
}
