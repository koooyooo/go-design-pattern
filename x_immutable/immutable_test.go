// [目的]
// Immutableパターンの目的は、インスタンスの状態を守ることです。
//
// [概要]
// Immutableのインスタンス状態の守り方は特殊です。そもそも更新を許さないのです。この仕組を物理的に用意します。
// 更新させなければ更新でおかしくなることもない。非常に頑固なイメージです。
//
// もちろん適用できる対象は限られます。例を挙げると Enumや ValueObjectです。
// これらは表現する対象こそ「種別」「値」と異なりますが、生成時に特定の値で状態を固定化し、以降変更を許さないという特徴があります。
// Enumに至っては型自体で全てを表現し、状態自体を持たないこともザラです。いずれにせよ状態が固定なら複数Threadから守るべき対象はありません。
//
// ValueObjectは、値自体を表現するものです。特徴としては生成時に値を固定し、その後は変更を許しません。
// 例えばJavaの java.awt.Color, java.lang.String, java.lang.Integer等は ValueObjectの一例です。
// Stringは連結 (ex. "Hello" + "World") で更新できるのでは、と思われがちですが、あれは新規生成された別インスタンスが返されています。
// また、Immutableとは逆に値を変更できるものを Mutableと呼びます。例えばJavaの java.lang.StringBuilderは Stringの Mutable表現です。
//
// Golangには 組み込みのWrapperクラスが無く、Enumもiota等の識別子で簡易的に表現することが多いためかあまり Immutableを見かけません。
// しかし作成し活用することは可能です。
//
// [作成法]
// 言語の力で制限を与えるという方針はSingletonに似ています。
// Immutableは、コンストラクタで値を代入し、その後は値の取得こそ許しますが、値の変更は許しません。
// また、値自体がImmutableでなければ、値の取得後に値自体を変更されてしまいますので、保持する状態はプリミティブ値か、それ自体もImmutableである必要があります。
package immutable_pattern

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/koooyooo/go-design-pattern/x_immutable/immutable_fragile"

	"github.com/koooyooo/go-design-pattern/x_immutable/immutable"
	"github.com/stretchr/testify/assert"
)

// 完成された Immutableのテストです
func TestImmutable(t *testing.T) {
	const LoopNum = 100

	var wg sync.WaitGroup
	wg.Add(100)

	score := 80
	s := immutable.NewScore(score)
	for i := 0; i < LoopNum; i++ {
		// 値の取得しか出来ないので同時アクセスされても何も出来ない
		// 下記のFragileと異なり、返却値もポインタではなく実値なので、変更が影響しない
		go func() {
			v := s.Value()
			// 返却値はコピー値に過ぎないので変更がscoreに影響しない
			v += 10
			// linterを落ち着かせるためだけの意味のないコード
			if v == 0 {
				fmt.Println("v == 0")
			}
			time.Sleep(200 * time.Millisecond)
			wg.Done()
		}()
	}
	wg.Wait()
	// 値は書き換えられていない
	assert.Equal(t, 80, s.Value())
}

func TestImmutableFragile(t *testing.T) {
	const LoopNum = 100

	var wg sync.WaitGroup
	wg.Add(LoopNum)

	score := 80
	s := immutable_fragile.NewScore(&score)
	for i := 0; i < LoopNum; i++ {
		// 取得した値が Immutableではないので値の書き換え可能
		// 同期制御も無いので、実行のたびに値が変わる
		go func() {
			v := s.Value()
			time.Sleep(200 * time.Millisecond)
			*v += 100
			wg.Done()
		}()
	}
	wg.Wait()
	// 値が書き換えられてしまう
	assert.NotEqual(t, 100, *s.Value())
	// 同期制御もないので、10080とも限らない (偶然10080になる場合もある)
	assert.NotEqual(t, 10080, *s.Value())
	fmt.Printf("Expected: 10080, Actual: %d \n", *s.Value())
}
