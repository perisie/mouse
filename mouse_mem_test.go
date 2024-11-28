package mouse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mouse_memory(t *testing.T) {
	var mouse Mouse

	mouse = Mouse_memory_new()

	_, err := mouse.Get("k1")

	assert.NotNil(t, err)

	_ = mouse.Put("k1", []byte("v1"))
	value, _ := mouse.Get("k1")

	assert.Equal(t, "v1", string(value))
}
