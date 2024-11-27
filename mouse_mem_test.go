package mouse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mouse_memory(t *testing.T) {
	mouse_memory := Mouse_memory_new()

	_, err := mouse_memory.Get("k1")

	assert.NotNil(t, err)

	_ = mouse_memory.Put("k1", []byte("v1"))
	value, _ := mouse_memory.Get("k1")

	assert.Equal(t, "v1", string(value))
}