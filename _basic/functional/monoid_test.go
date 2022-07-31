package functional

import (
	"github.com/koooyooo/go-design-pattern/_basic/functional/collection"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringMonoid(t *testing.T) {
	var s string
	s = stringMonoid.FoldRight([]string{"a", "b", "c"})
	assert.Equal(t, "abc", s)

	s = stringMonoid.FoldLeft([]string{"a", "b", "c"})
	assert.Equal(t, "abc", s)
}

func TestIntPlusMonoid(t *testing.T) {
	var i int
	i = intPlusMonoid.FoldRight([]int{1, 2, 3})
	assert.Equal(t, 6, i)

	i = intPlusMonoid.FoldLeft([]int{1, 2, 3})
	assert.Equal(t, 6, i)
}

func TestListMonoid(t *testing.T) {
	var l collection.List[int]
	l = listMonoid[int]().FoldRight(collection.List[collection.List[int]]{{1, 2, 3}, {4, 5, 6}})
	assert.Equal(t, collection.List[int]{1, 2, 3, 4, 5, 6}, l)
}
