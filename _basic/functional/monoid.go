package functional

type Monoid[T any] struct {
	zero T
	op   func(T, T) T
}

//// List[A]
//def foldRight[B](zero: B)(f: (A, B) => B): B =
//if (this.isEmpty) zero
//else f(head, tail.foldRight(zero)(f))

func (m Monoid[T]) FoldLeft(t list[T]) T {
	if len(t) == 0 {
		return m.zero
	}
	ln := len(t)
	return m.op(m.FoldLeft(t[0:ln-1]), t[ln-1])
}

func (m Monoid[T]) FoldRight(t list[T]) T {
	if len(t) == 0 {
		return m.zero
	}
	return m.op(t[0], m.FoldRight(t[1:]))
}

func (m Monoid[T]) Fold(t list[T]) T { return m.FoldLeft(t) }

var stringMonoid = Monoid[string]{
	zero: "",
	op: func(a, b string) string {
		return a + b
	},
}

var intPlusMonoid = Monoid[int]{
	zero: 0,
	op: func(a, b int) int {
		return a + b
	},
}

var intMultiplyMonoid = Monoid[int]{
	zero: 1,
	op: func(a, b int) int {
		return a * b
	},
}

func listMonoid[T any]() Monoid[list[T]] {
	return Monoid[list[T]]{
		zero: nil,
		op: func(a, b list[T]) list[T] {
			return append(a, b...)
		},
	}
}
