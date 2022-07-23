// Package iterator_pattern
//
// [目的]
// Iteratorの目的は、スライスの要素を安全に取り出すことです。
//
// [概要]
// 通常のIndexを用いた取り出しは、スライスの要素数を超えパニックとなる可能性があります。
// Iteratorを用いると、スライスの要素は内部管理され、この様な危険がありません。
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
package iterator_pattern

import (
	"github.com/koooyooo/go-design-pattern/iterator/iterator"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestIterator(t *testing.T) {
	tests := []struct {
		name     string
		origin   []int
		expected []int
	}{
		{
			name:     "1 to 5",
			origin:   []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "random",
			origin:   []int{-1, 2, 0, -4},
			expected: []int{-1, 2, 0, -4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Iterator を生成する
			ite := iterator.NewSliceIterator[int](test.origin)
			var actual []int
			// HasNext()と Next()を組み合わせることで indexアクセスせずにループ処理を実現する
			for ite.HasNext() {
				actual = append(actual, ite.Next())
			}
			// 元の要素が正しくコピーされていることを保証する。
			assert.Equal(t, test.expected, actual)
		})
	}
}
