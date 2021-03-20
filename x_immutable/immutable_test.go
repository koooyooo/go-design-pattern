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
	"testing"

	"github.com/koooyooo/go-design-pattern/x_immutable/immutable"
	"github.com/stretchr/testify/assert"
)

func TestImmutable(t *testing.T) {
	s := immutable.NewScore(100)
	for i := 0; i < 100; i++ {
		// 値の取得しか出来ないので同時アクセスされても何も出来ない
		go func() {
			s.Value()
		}()
	}
	assert.Equal(t, 100, s.Value())
}
