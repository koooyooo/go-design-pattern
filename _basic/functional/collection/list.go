package collection

type List[T any] []T

func (l *List[T]) Contains(t T) bool {
	return contains[T](*l, t)
}

func (l *List[T]) ContainsAll(t ...T) bool {
	return contains[T](*l, t...)
}

func (l *List[T]) Filter(f func(t T) bool) List[T] {
	return filter[T](*l, f)
}

func (l *List[T]) Map(f func(t T) T) List[T] {
	return map1[T, T](*l, f)
}

func (l *List[T]) FoldLeft(zero T, f func(t1, t2 T) T) T {
	return foldLeft[T](*l, zero, f)
}

func (l *List[T]) FoldRight(zero T, f func(t1, t2 T) T) T {
	return foldRight[T](*l, zero, f)
}

func (l *List[T]) Reverse() List[T] {
	var tmp List[T] = reverse(*l)
	return tmp
}

func (l *List[T]) HeadTail() (*T, List[T], bool) {
	return headTail(*l)
}
