package option

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {
	var o, n Option[int]

	// Some
	o = NewSome[int](1)
	assert.False(t, o.IsEmpty())
	v, ok := o.Get()
	assert.True(t, ok)
	assert.Equal(t, 1, v)
	assert.Equal(t, 1, o.GetOrElse(2))

	// None
	n = NewNone[int]()
	assert.True(t, n.IsEmpty())
	v, ok = n.Get()
	assert.False(t, ok)
	assert.Nil(t, v)
	assert.Equal(t, 2, n.GetOrElse(2))
}
