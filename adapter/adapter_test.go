// [目的]
// Adapterパターンの目的は、旧インターフェイスを活用しながら新インターフェイスを利用することです。
//
// [概要]
// とある機能の新旧インターフェイスには形式に差があって現在利用できなくあんっているとします。
// ただ、形式に差があったとしても、本質的な情報量に差がなければ、間に変換用クラスを差し込めば利用可能な筈です。
// この変換用クラスがAdapterです。
//
// Adapterを活用すれば 新旧のインターフェイス差を埋めることができます。さらには次の長所も発生します。
// - Adapterに変換処理が凝縮されるため、ロジックにおける新旧変換の表現性が高い
// - 新旧のモジュールの更新が最低限（もしくは皆無）となり、修正コストが少なく実現速度も高い
// - OpenClosed原則の Closedを Adapterで実現できる
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

	// 本来Legacyな APIとして稼働しているが
	s, err := l.Receive("get", "/path", nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, "legacy: get /path", s)

	// Adapterを介することでModernAPIに代入し
	var m adapter.ModernAPI = adapter.Adapter{
		Legacy: l,
	}

	// 機能仕様も満たすことが可能 (内部でLegacyAPIの実装に委譲)
	s, err = m.Receive(httptest.NewRequest("get", "http://sample.com/path", nil))
	assert.NoError(t, err)
	assert.Equal(t, "legacy: get /path", s)
}
