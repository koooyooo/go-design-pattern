package function

import (
	"github.com/google/go-cmp/cmp"
)

type list[T any] []T

func (l *list[T]) Contains(t T) bool {
	for _, v := range *l {
		if cmp.Equal(v, t) {
			return true
		}
	}
	return false
}

func (l *list[T]) Map(f func(t T) T) list[T] {
	var result list[T]
	for _, v := range *l {
		result = append(result, f(v))
	}
	return result
}
