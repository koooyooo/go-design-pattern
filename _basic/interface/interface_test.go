package _interface

import (
	"github.com/koooyooo/go-design-pattern/_basic/interface/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 様々な履行方法があるが、全てInterfaceを満たせることを確認
func TestInterface(t *testing.T) {
	impls := []any{
		&model.InterfaceImpl{},
		model.InterfaceImplVal{},
		&model.InterfaceImplVal{},
		&model.InterfaceImplEmbed{},
		model.InterfaceImplEmbedVal{},
		&model.InterfaceImplEmbedVal{},
	}
	for _, impl := range impls {
		var i, ok = impl.(model.Interface)
		assert.True(t, ok)
		assert.Equal(t, "do something", i.Do())
	}
}
