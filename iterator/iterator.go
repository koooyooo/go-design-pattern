// Package iterator_pattern
//
// [目的]
// Iteratorの目的は、スライスの要素を安全に取り出すことです。
//
// [概要]
// 通常のIndexを用いた取り出しは、スライスの要素数を超えパニックとなる可能性があります。
// Iteratorを用いると、スライスの要素は内部管理され、この様な危険がありません。
//
// Go言語の場合は rangeを用いたループにより、indexによる要素アクセス回避可能なので Iteratorの役割は満たせていると言えます。
//
// [危うい要素アクセス]
// for i := 0; i < len(array); i++ {
//     fmt.Println(array[i])
// }
//
// [安全な要素アクセス: range]
// for _, a := range array {
//     fmt.Println(a)
// }
//
// [安全な要素アクセス: Iterator]
// for ite.HasNext() {
//    fmt.Println(ite.Next())
// }
//
// Iteratorを用いれば、言語機能に頼らずに様々なバリエーションでこれを実現することができます。
// 例えば、ファイルから一行ずつ読み込む Iterator, DBから一レコードずつ読み込む Iterator, APIから1レスポンスずつ読み込む Iterator 等です。
//
package iterator_pattern

// Iterator は、次の要素があるかどうかを返すイテレータインターフェイス
type Iterator[V any] interface {
	HasNext() bool
	Next() V
}

// NewSliceIterator は、スライスをベースとしたイテレータを生成する
func NewSliceIterator[V any](vs []V) Iterator[V] {
	i := &sliceIterator[V]{
		vs: vs,
	}
	return i
}

type sliceIterator[V any] struct {
	current int
	vs      []V
}

func (i *sliceIterator[V]) HasNext() bool {
	return i.current < len(i.vs)
}

func (i *sliceIterator[V]) Next() V {
	v := i.vs[i.current]
	i.current++
	return v
}
