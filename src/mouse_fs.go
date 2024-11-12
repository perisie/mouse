package mouse

import (
	"fmt"
	"io"
	"os"
)

type Mouse_fs struct {
	data_dir_path string
}

func (m *Mouse_fs) Put(key string, value []byte) error {
	file_path := m.data_file_path(key)
	f, err := os.Create(file_path)
	if err != nil {
		return err
	}
	_, err = f.Write(value)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}

func (m *Mouse_fs) Get(key string) ([]byte, error) {
	file_path := m.data_file_path(key)
	f, err := os.Open(file_path)
	if err != nil {
		return nil, err
	}
	value, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (m *Mouse_fs) data_file_path(key string) string {
	return fmt.Sprintf("%v/%v.dt", m.data_dir_path, m.normalize_key(key))
}

func (m *Mouse_fs) normalize_key(key string) string {
	key_n := ""
	for _, k := range key {
		for _, ch := range "abcdefghijklmnopqrstuvwxyz0123456789_" {
			if k == ch {
				key_n += string(ch)
			}
		}
	}
	return key_n
}

func Mouse_fs_new(data_dir_path string) *Mouse_fs {
	_ = os.Mkdir(data_dir_path, 0755)
	return &Mouse_fs{data_dir_path: data_dir_path}
}
