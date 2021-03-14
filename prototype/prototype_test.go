// [目的]
// Prototypeパターンの目的はインスタンスの複製の省力化です
//
// [概要]
// 特に内部で保持している状態が多い、または複雑なインスタンスの場合、状態を複製するのが大変です。
// しかし手書きでの複製が大変でもコピーが楽な局面は多いものです。Prototypeはコピーしてインスタンスを生成することにあります。
//
// [実現]
// Javaの場合は全てのインスタンスはポインタなので、言語特有の Object#clone()をコールします。
// Golangの場合はstructがポインタとは限らないので、参照先の struct値を別変数にコピーすることで実現可能です。
package prototype_pattern

import (
	"testing"

	"github.com/koooyooo/go-design-pattern/prototype/clonable"

	"github.com/stretchr/testify/assert"
)

func TestPrototype(t *testing.T) {
	// 構造体"Android"を生成し状態を作り込む
	a := clonable.Android{
		Name: "Droid",
		Log:  []string{"Walk", "Sleep", "Sleep Twice"},
	}
	// Cloneを作成することで状態の複製を簡易化
	b := a.Clone()
	assert.Equal(t, "Droid", b.Name)
	assert.ElementsMatch(t, []string{"Walk", "Sleep", "Sleep Twice"}, b.Log)

	// 片系統の値を変更してももう一方に影響が無いことを確認
	b.Name = "AnDroid"
	assert.Equal(t, "Droid", a.Name)
	assert.Equal(t, "AnDroid", b.Name)
}
