package extension

import (
	"github.com/koooyooo/go-design-pattern/_basic/extension/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 親の型に子を代入することができない
func TestAssignChildToParent(t *testing.T) {
	// 値の代入
	//var m1 extension.Mammal = extension.Human{} // error

	// 参照の代入
	//var m2 *extension.Mammal = &extension.Human{} // error
}

// Interfaceを介在させれば代入は可能
func TestAssignChildToIF(t *testing.T) {
	type mammalIF interface {
		Do() string
	}
	// 値の代入
	var _ mammalIF = model.Human{}

	// 参照の代入
	var _ mammalIF = &model.Human{}
}

func TestOverride(t *testing.T) {
	// 値の代入
	var m1 = model.Human{}
	assert.Equal(t, "do something by human", m1.Do())

	// 参照の代入
	var m2 = &model.Human{}
	assert.Equal(t, "do something by human", m2.Do())
}
