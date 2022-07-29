package functional

import "github.com/google/go-cmp/cmp"

type Option[T any] interface {
	IsEmpty() bool
	Get() (*T, bool)
	GetOrElse(t T) T
}

type Some[T any] struct {
	value T
}

func NewSome[T any](v T) Option[T] {
	return &Some[T]{value: v}
}

func (s *Some[T]) IsEmpty() bool {
	return false
}

func (s *Some[T]) Get() (*T, bool) {
	return &s.value, true
}

func (s *Some[T]) GetOrElse(t T) T {
	if cmp.Equal(s, nil) {
		return t
	}
	return s.value
}

type None[T any] struct{}

func NewNone[T any]() Option[T] {
	return &None[T]{}
}

func (n *None[T]) IsEmpty() bool {
	return true
}

func (n *None[T]) Get() (*T, bool) {
	return nil, false
}

func (n *None[T]) GetOrElse(t T) T {
	return t
}
