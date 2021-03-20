// [目的]
// SingleThreadExecutionの目的は、複数Threadの同時実行による不整合を抑止することです。
// 何の不整合かというと、主に対象のメソッドが更新する状態の不整合となります。
//
// [概要]
// 例えば銀行口座の取り扱いを例に取ります。銀行口座には残高(Balance)という内部状態があります。
// この残高を1件1件の更新リクエスト毎に順番に更新してゆけば、問題は起こらないのですが、これが一気に更新すると事情が異なります。
// 例えば、100人の銀行員が同じ瞬間に同じ口座を更新すると、処理途中でまだ最終確定していない残高に対して計算処理をしてしまったり、
// 同じ残高を同時編集しようとして、編集が上手く反映されなかったりします。
// この様な問題は、僅かなブレが許される傾向調査などであれば問題ないのですが、お金のやり取りでは許容できません。
//
// 解決策の一つが SingleThreadExecution, つまり整合性を担保するフローにおいては1度に1-Threadのみ通過を許すことです。
// 複数人からの書き込みで問題が起きるのであれば、複数人からの書き込みを禁止してしまおうという発想です。
// 昔の銀行では、複数銀行員が同じ口座台帳に対して同時に書き込むのは困難だったと思いますが、これも意図せず実装された SingleThreadExecutionであるわけです。
//
// [実現方法]
// Javaでは Methodに synchronized修飾子を付けたり、synchronizedブロックを宣言したりしてSingleThreadの区間を宣言しました。
// Goでは sync.Mutexの Lock()と Unlock()により同等のスコープを宣言します。
// 「整合性を担保すべき一連の状態更新はそのスコープの中で完結させる」という点は一緒です。
//
// [シナリオ]
// 残高 1000の銀行口座に対し、同時に100の処理Threadで更新をかけます。
// 1回の更新内容は +100と -100という一連の操作です。理論的には最終的な残高は変化しません。
// しかし、SingleThreadExecutionを実装していないと、高い確率で不整合が発生してしまいます。
//
// これは Thread-Aが更新途中の値を元にThread-Bが計算してしまうというレースコンディションに起因するものや、
// メモリ上の同じ値を同時編集することに起因するものがあります。
//
// そこで、一連のTransactionを表現するスコープにおいて、SingleThreadExecutionを実装して不整合を防いでいます。
package single_thread_execution_pattern

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/koooyooo/go-design-pattern/x_single_thread_execution/un_synchronized"

	"github.com/koooyooo/go-design-pattern/x_single_thread_execution/synchronized"
)

// SingleThreadExecutionを実装していると不整合はない
func TestSingleThreadExecution(t *testing.T) {
	const numTx = 100

	// 処理数管理用の WaitGroupを用意（テスト都合)
	var w sync.WaitGroup
	w.Add(numTx)

	// 残高1000の口座に対し 100Threadから同時に+100,-100の取引を実施
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

// SingleThreadExecutionを実装していないと高い確率で不整合が発生する
func TestNonSingleThreadExecution(t *testing.T) {
	const numTx = 100

	// 処理数管理用の WaitGroupを用意（テスト都合)
	var w sync.WaitGroup
	w.Add(numTx)

	// 残高1000の口座に対し 100Threadから同時に+100,-100の取引を実施
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
