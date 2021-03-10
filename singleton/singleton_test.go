package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/koooyooo/go-design-pattern/singleton/singleton"
)

// Singletonは唯一のインスタンスを用意し、それ以外を生成できなくするパターンです。
// そのため、唯一のインスタンスしか存在しないことを保証できます。
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
