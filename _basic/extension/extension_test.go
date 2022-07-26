// Package extension
//
// Go言語には継承がないので Embedを用いる。
// Embedは機能的は提供するが、型は提供しない。従ってJava等の以下の機能は期待できない。
// - 代入を期待する「型」としての親クラス
// - 継承を期待する「型」としての抽象クラス
//
// Go言語での型はあくまでインターフェイスが提供し、Embedは機能拡張するにとどまる。
// Javaでも型が重要になった場合、専用のインターフェイスを用意するのは良いプラクティスなので、それが文法的に強制されている事にデメリットはない。
// インターフェイスで型を提供し、Embedで機能を提供するというのが基本パターンとなる。
// Embedの場合、単なる委譲と異なり委譲元が委譲先と同等のメソッドを用意し委譲先に仲介する処理を書く必要もないので、実装上の負担も少ない。
//
// 但し、継承を前提とする Template Methodの様なパターンには工夫が必要になるので、そこは意識しなくてはならない。
//
package extension

import (
	"github.com/koooyooo/go-design-pattern/_basic/extension/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Embedベースの拡張では、親クラスに当たる構造体への代入ができない（そもそも親クラスではない）
// - 親クラスの「型」としての役割は期待できない
// - 抽象クラスという存在もない
// - リスコフの置換原則は適用できない（そもそも親として振る舞えないので置換も不可能）
// - 「型」が必要なら「インターフェイス」を使う
func TestAssign(t *testing.T) {
	// 親の型に子を代入することができない
	t.Run("parent", func(t *testing.T) {
		//var m1 extension.Mammal = extension.Human{} // error
		//var m2 *extension.Mammal = &extension.Human{} // error
	})

	// Interfaceを介在させれば代入は可能
	t.Run("interface", func(t *testing.T) {
		type mammalIF interface {
			Do() string
		}
		var _ mammalIF = model.Human{}
		var _ mammalIF = &model.Human{}
	})
}

// メソッドのオーバーライドは可能
func TestOverride(t *testing.T) {
	// 親のメソッドを
	t.Run("parent", func(t *testing.T) {
		assert.Equal(t, "do something", model.Mammal{}.Do())
		assert.Equal(t, "do something", (&model.Mammal{}).Do())
	})

	// 子でオーバーライド可能
	t.Run("child", func(t *testing.T) {
		assert.Equal(t, "do something by human", model.Human{}.Do())
		assert.Equal(t, "do something by human", (&model.Human{}).Do())
	})
}
