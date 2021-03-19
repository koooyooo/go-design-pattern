// [目的]
// SingleThreadExecutionの目的は、複数Threadの同時実行による内部状態の不整合を抑止することです。
// このパターンを用いないと、対象スコープに複数Threadが入り込みます。
//
//
package x_single_thread_execution_pattern

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/koooyooo/go-design-pattern/x_single_thread_execution/un_synchronized"

	"github.com/koooyooo/go-design-pattern/x_single_thread_execution/synchronized"
)

func TestSingleThreadExecution(t *testing.T) {
	var w sync.WaitGroup
	w.Add(100)
	s := synchronized.Account{}
	// 100 ThreadからTransaction処理を実施
	for i := 0; i < 100; i++ {
		go s.Transaction(&w)
	}
	w.Wait()

	// Plus,Minus のログが必ず交互に出力されることを確認
	assert.NotContains(t, "Plus,Plus", s.String())
	assert.NotContains(t, "Minus,Minus", s.String())

	// 値を保証できていることを確認
	assert.Equal(t, 0, s.Amount())
}

func TestNonSingleThreadExecution(t *testing.T) {
	var w sync.WaitGroup
	w.Add(100)
	s := un_synchronized.Account{}
	// 100 ThreadからTransaction処理を実施
	for i := 0; i < 100; i++ {
		go s.Transaction(&w)
	}
	w.Wait()

	// Plus,Minus のログが必ずしも交互に出力されないことを確認
	assert.Contains(t, s.String(), "Plus,Plus")
	assert.Contains(t, s.String(), "Minus,Minus")

	// 値を保証できないことを確認
	assert.NotEqual(t, 0, s.Amount())
}
