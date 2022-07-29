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

func (l *list[T]) FoldLeft(zero T, f func(t1, t2 T) T) T {
	var tmp = zero
	for _, v := range *l {
		tmp = f(tmp, v)
	}
	return tmp
}

func (l *list[T]) FoldRight(zero T, f func(t1, t2 T) T) T {
	return l.Reverse().FoldLeft(zero, f)
}

func (l *list[T]) Reverse() *list[T] {
	var tmp = *l
	for i := 0; i < len(*l)/2; i++ {
		tmp[i], tmp[len(tmp)-1-i] = tmp[len(tmp)-1-i], tmp[i]
	}
	return &tmp
}
