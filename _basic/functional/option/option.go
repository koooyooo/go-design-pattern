package option

import (
	"github.com/google/go-cmp/cmp"
	"github.com/koooyooo/go-design-pattern/_basic/functional/container"
)

type Option[T any] interface {
	container.Container[T]
	IsEmpty() bool
	GetOrElse(t T) T
}

type Some[T any] struct {
	value T
}

func NewSome[T any](v T) Option[T] {
	return &Some[T]{value: v}
}

func (s *Some[T]) Unit(v T) container.Container[T] {
	return &Some[T]{value: v}
}

func (s *Some[T]) IsEmpty() bool {
	return false
}

func (s *Some[T]) Get() (*T, bool) {
	return &(s.value), true
}

func (s *Some[T]) GetOrElse(t T) T {
	if cmp.Equal(s, nil) {
		return t
	}
	return s.value
}
func (s *Some[T]) Map2(a, b container.Container[any], f func(a, b any) any) container.Container[any] {
	va, _ := a.Get()
	vb, _ := b.Get()
	vc := f(va, vb)
	return &Some[any]{value: vc}
}

type None[T any] struct{}

func NewNone[T any]() Option[T] {
	return &None[T]{}
}

func (s *None[T]) Unit(v T) container.Container[T] {
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

func (_ *None[T]) Map2(a, b container.Container[any], f func(a, b any) any) container.Container[any] {
	return &None[any]{}
}
