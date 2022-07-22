// Package iterator
//
// [目的]
// Iteratorの目的は、スライスの要素を安全に取り出すことです。
//
// [概要]
// 通常のIndexを用いた取り出しは、スライスの要素数を超えパニックとなる可能性があります。
// Iteratorを用いると、スライスの要素は内部管理され、この様な危険がありません。
//
//
// [危うい要素アクセス]
// for i := 0; i < len(array); i++ {
//     fmt.Println(array[i])
// }
//
// [安全な要素アクセス: Iterator]
// for ite.HasNext() {
//    fmt.Println(ite.Next())
// }
//
// なお、Go言語の場合は rangeにより indexによる要素アクセスを回避可能なので Iteratorの役割は満たせていると言えます。
//
// [安全な要素アクセス: range]
// for _, a := range array {
//     fmt.Println(a)
// }
//
// 但し、rangeを利用できるのは 配列・スライス・マップに限定されます。
// Iteratorを使えば、言語機能に頼らずに様々なバリエーションで 安全な要素アクセスを実現することができます。
// 例えば、ファイルから一行ずつ読み込む Iterator, DBから一レコードずつ読み込む Iterator, APIから1レスポンスずつ読み込む Iterator 等です。
//
package iterator

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
