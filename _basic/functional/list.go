package functional

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

func (l *list[T]) Filter(f func(t T) bool) list[T] {
	var result list[T]
	for _, v := range *l {
		if !f(v) {
			continue
		}
		result = append(result, v)
	}
	return result
}

func (l *list[T]) Map(f func(t T) T) list[T] {
	var result = make(list[T], len(*l))
	for i, v := range *l {
		result[i] = f(v)
	}
	return result
}
