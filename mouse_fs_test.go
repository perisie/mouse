package mouse

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_mouse_fs(t *testing.T) {
	mouse := get_mouse_fs()

	v, _ := mouse.Get("k1")

	assert.Equal(t, "", string(v))

	_ = mouse.Put("k1", []byte("v1"))
	v, _ = mouse.Get("k1")

	assert.Equal(t, "v1", string(v))
}

func Test_mouse_fs_parallel_write(t *testing.T) {
	mouse := get_mouse_fs()
	values := map[string]bool{}

	for i := 0; i < 1000; i++ {
		v := "data" + fmt.Sprint(i)
		go func() {
			_ = mouse.Put("k", []byte(v))
		}()
		values[v] = true
	}

	v_str := ""

	for v_str == "" {
		v, _ := mouse.Get("k")
		v_str = string(v)
	}

	ok := false

	for v_want := range values {
		fmt.Printf("%v\t%v\n", v_str, v_want)
		if v_want == v_str {
			ok = true
		}
	}

	assert.True(t, ok)
}

func get_mouse_fs() Mouse {
	var mouse Mouse
	_ = os.RemoveAll("data")
	mouse = Mouse_fs_new("data")
	return mouse
}
