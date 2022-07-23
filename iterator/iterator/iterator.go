package iterator

// Iterator は、次の要素があるかどうかを返すイテレータインターフェイス
type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

// NewSliceIterator は、スライスをベースとしたイテレータを生成する
func NewSliceIterator[T any](vs []T) Iterator[T] {
	i := &sliceIterator[T]{
		slice: vs,
	}
	return i
}

type sliceIterator[T any] struct {
	idx   int
	slice []T
}

func (i *sliceIterator[T]) HasNext() bool {
	return i.idx < len(i.slice)
}

func (i *sliceIterator[T]) Next() T {
	v := i.slice[i.idx]
	i.idx++
	return v
}
