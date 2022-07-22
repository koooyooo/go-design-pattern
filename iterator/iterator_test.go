package iterator_pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator(t *testing.T) {
	tests := []struct {
		name     string
		origin   []int
		expected []int
	}{
		{
			name:     "1 to 5",
			origin:   []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "random",
			origin:   []int{-1, 2, 0, -4},
			expected: []int{-1, 2, 0, -4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ite := NewSliceIterator[int](test.origin)
			var actual []int
			for ite.HasNext() {
				actual = append(actual, ite.Next())
			}
			assert.Equal(t, test.expected, actual)
		})
	}
}
