package functional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	var l list[int] = []int{1, 2, 3}

	assert.True(t, l.Contains(2))
	assert.False(t, l.Contains(4))
}

func TestFilter(t *testing.T) {
	var l list[int] = []int{1, 2, 3}

	tests := []struct {
		f    func(i int) bool
		want list[int]
	}{
		{
			f:    func(i int) bool { return i%2 == 0 },
			want: list[int]{2},
		}, {
			f:    func(i int) bool { return i%2 == 1 },
			want: list[int]{1, 3},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, l.Filter(test.f))
	}
}

func TestMap(t *testing.T) {
	var l list[int] = []int{1, 2, 3}

	var doubled = l.Map(func(i int) int {
		return i * 2
	})

	assert.Equal(t, list[int]{2, 4, 6}, doubled)
}
