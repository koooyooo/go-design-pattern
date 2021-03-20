// [目的]
// SingleThreadExecutionの目的は、複数Threadの同時実行による内部状態の不整合を抑止することです。
// このパターンを用いないと、対象スコープに複数Threadが入り込みます。
//
//
package x_single_thread_execution_pattern

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/koooyooo/go-design-pattern/x_single_thread_execution/un_synchronized"

	"github.com/koooyooo/go-design-pattern/x_single_thread_execution/synchronized"
)

func TestSingleThreadExecution(t *testing.T) {
	const numTx = 100

	// 処理数管理用の WaitGroupを用意（テスト都合)
	var w sync.WaitGroup
	w.Add(numTx)

	// 残高1000の口座に対し非同期で100回 +100 -100の取引を実施
	a := synchronized.NewAccount(&w, 1000)
	for i := 0; i < numTx; i++ {
		go a.PlusMinusTransaction(100, 100)
	}
	w.Wait()

	// Plus,Minus のログが必ず交互に出力されることを確認
	assert.NotContains(t, "Plus,Plus", a.String())
	assert.NotContains(t, "Minus,Minus", a.String())

	// 値を保持できていることを確認
	assert.Equal(t, 1000, a.Amount())
	fmt.Printf("Single: %b\n", a.Amount())
}

func TestNonSingleThreadExecution(t *testing.T) {
	const numTx = 100

	// 処理数管理用の WaitGroupを用意（テスト都合)
	var w sync.WaitGroup
	w.Add(numTx)

	// 残高1000の口座に対し非同期で100回 +100 -100の取引を実施
	a := un_synchronized.NewAccount(&w, 1000)
	for i := 0; i < numTx; i++ {
		go a.PlusMinusTransaction(100, 100)
	}
	w.Wait()

	// Plus,Minus のログが必ずしも交互に出力されないことを確認 (100%交互出力が崩れる訳ではない)
	assert.Contains(t, a.String(), "Plus,Plus")
	assert.Contains(t, a.String(), "Minus,Minus")

	// 値を保持できないことを確認 (100%保持できない訳ではない)
	assert.NotEqual(t, 1000, a.Amount())
	fmt.Printf("NonSingle: %b\n", a.Amount())
}
