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
		model.InterfaceImplRaw{},
		&model.InterfaceImplRaw{},
		&model.InterfaceImplEmbed{},
		model.InterfaceImplEmbedRaw{},
		&model.InterfaceImplEmbedRaw{},
	}
	for _, impl := range impls {
		var i, ok = impl.(model.Interface)
		assert.True(t, ok)
		assert.Equal(t, "do something", i.Do())
	}
}
