package functional

type list[T any] []T

func NewList[T any](t ...T) list[T] {
	return t
}

func (l *list[T]) Contains(t T) bool {
	return Contains[T](*l, t)
}

func (l *list[T]) Filter(f func(t T) bool) list[T] {
	return Filter[T](*l, f)
}

func (l *list[T]) Map(f func(t T) T) list[T] {
	return Map[T, T](*l, f)
}

func (l *list[T]) FoldLeft(zero T, f func(t1, t2 T) T) T {
	return FoldLeft[T](*l, zero, f)
}

func (l *list[T]) FoldRight(zero T, f func(t1, t2 T) T) T {
	return FoldRight[T](*l, zero, f)
}

func (l *list[T]) Reverse() *list[T] {
	var tmp list[T] = Reverse(*l)
	return &tmp
}
