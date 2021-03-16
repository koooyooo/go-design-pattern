// [目的]
// Adapterパターンの目的は、旧インターフェイスを活用しながら新インターフェイスを利用することです。
//
// [概要]
// 新旧インターフェイスの入出力に形式の差こそあれ情報量の差がなければ、間に変換用クラスを差し込むだけで利用可能な筈です。
// この変換用クラスがAdapterです。Adapterを差し込むことで 新旧のインターフェイス差を最小限の実装で埋めることができます。
// また、Adapterのコードには新旧インターフェイスの変換処理が集約されるため、新旧属性の対応を表現するコードとしても重宝します。
//
// [シナリオ]
// 今回のシナリオは HTTPのリクエストを受け付ける APIをレガシーなインターフェイスからモダンなインターフェイスに移行する局面です。
// HTTPのリクエストを受け付けるという性質は変わらないので、Adapterの変換による互換性担保が可能になります。
// また、Adapter自体がモダンなインターフェイスを満足させている部分もポイントです。
package adapter_pattern

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/koooyooo/go-design-pattern/adapter/adapter"
)

func TestAdapter(t *testing.T) {
	// 実体はLegacyAPI
	var l adapter.LegacyAPI = &adapter.LegacyAPIImpl{}

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
