package function

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	var l list[int] = []int{1, 2, 3}

	assert.True(t, l.Contains(2))
	assert.False(t, l.Contains(4))
}

func TestMap(t *testing.T) {
	var l list[int] = []int{1, 2, 3}

	var doubled = l.Map(func(i int) int {
		return i * 2
	})

	assert.Equal(t, list[int]{2, 4, 6}, doubled)
}
