package prototype_pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Android struct {
	Name string
}

func TestPrototype(t *testing.T) {
	a := Android{
		Name: "Droid",
	}
	// 非ポインタの参照を渡すだけでクローン可能
	b := a
	assert.Equal(t, "Droid", b.Name)

	b.Name = "AnDroid"
	assert.Equal(t, "Droid", a.Name)
	assert.Equal(t, "AnDroid", b.Name)
}
