package extension

import (
	"testing"

	"github.com/koooyooo/go-design-pattern/0_extension/extension"
	"github.com/stretchr/testify/assert"
)

// 親の型に子を代入することができない
func TestAssignChildToParent(t *testing.T) {
	//var m1 extension.Mammal
	//m1 = extension.Human{} // error

	//var m2 *extension.Mammal
	//m2 = &extension.Human{} // error
}

// 直接親の型に代入ではなく、同等のInterfaceを用意しそこに代入することは可能
func TestDo(t *testing.T) {
	type MurmalIF interface {
		Do() string
	}
	var m MurmalIF
	m = extension.Human{}
	assert.Equal(t, "do something by human", m.Do())
}
