package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/koooyooo/go-design-pattern/singleton/singleton"
)

// Singleton は唯一のインスタンスを生成するパターンです。具体的には以下の特徴でこれを実現します。
// - 外部から直接生成できなくし、替わりに参照取得用の関数を用意する
// - その関数から何度取得しても同じ参照が返る
// Singleton を適用することで、次の様なことが実現できます。
// - Enumの様に structで種別を表現する
// - Flyweightパターンとして Instance生成の負荷を抑えて軽量化する
// また、DIコンテナの中ではこの様なコード的な制約は持たせませんが Serviceや Repository等のコンポーネントを1つだけ生成します。
// この様に、デザインパターンとして適用するSingletonだけではなく、事実上Singletonが成立する場合もあり、シンプルな分より理想的です。
// Singletonはクラス(struct)やインターフェイス同士の組み合わせによるパターンというよりは、struct内部のパターンとなります。
func TestSingleton(t *testing.T) {
	// Instance()以外で生成できない => singleton.singleton{} は不可能
	s1 := singleton.Instance()
	s2 := singleton.Instance()

	// インスタンスの内容が同じだけでなく
	assert.Equal(t, s1, s2)
	// インスタンスの実体も同じ
	assert.True(t, s1 == s2)
}

// 逆に一般的な生成パターンは、コンストラクタ関数の中でインスタンスを生成します。
// そのため、同じ値を持ちますが(同値)、参照は同じ(同一)ではありません。
func TestNonSingleton(t *testing.T) {
	s1 := singleton.NewNormal()
	s2 := singleton.NewNormal()

	// インスタンスの内容は同じだが
	assert.Equal(t, s1, s2)
	// インスタンスの実体は異なる
	assert.False(t, s1 == s2)
}
