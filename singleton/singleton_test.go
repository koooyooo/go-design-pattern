package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/koooyooo/go-design-pattern/singleton/singleton"
)

func TestSingleton(t *testing.T) {
	// Instance()以外で生成できない
	// => singleton.singleton{} は不可能
	s1 := singleton.Instance()
	s2 := singleton.Instance()
	s3 := singleton.Instance()

	// どのインスタンスも同じ参照
	assert.Equal(t, s1, s2)
	assert.Equal(t, s1, s3)
	assert.Equal(t, s2, s3)
}
