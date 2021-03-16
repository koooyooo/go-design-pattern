// [目的]
// Adapterパターンの目的は、レガシーのインターフェイスを活用しながら
//
// [概要]
//
package adapter_pattern

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/koooyooo/go-design-pattern/adapter/adapter"
)

func TestAdapter(t *testing.T) {
	// 実体はLegacyAPI
	l := adapter.NewLegacyAPI()

	// Adapterを介することでModernAPIとして適用可能
	var m adapter.ModernAPI = adapter.Adapter{
		Legacy: l,
	}

	// 機能的にもLegacyAPIを活用することで問題なく稼働
	r := httptest.NewRequest("get", "http://sample.com/path", nil)
	s, err := m.Receive(r)

	assert.NoError(t, err)
	assert.Equal(t, "legacy: get /path", s)
}
