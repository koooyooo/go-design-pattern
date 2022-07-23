package adapter

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	// LegacyAPI はレガシーなAPI (httpパッケージを活用せずに実装)
	LegacyAPI interface {
		Receive(method, url string, header map[string][]string, body []byte) (string, error)
	}
	// ModernAPI はモダンなAPI (httpパッケージを活用して実装)
	ModernAPI interface {
		Receive(r *http.Request) (string, error)
	}
)

// LegacyAPIImpl はレガシーAPI実装
type LegacyAPIImpl struct{}

// Receive はレガシーAPIの実装
func (l LegacyAPIImpl) Receive(method, url string, header map[string][]string, body []byte) (string, error) {
	return fmt.Sprintf("legacy: %s %s", method, url), nil
}

// Adapter はアダプタの実装
type Adapter struct {
	// 委譲対象のLegacyAPIを保持
	Legacy LegacyAPI
}

// Receive は変換処理の実装でありモダンAPIのインターフェイスも満足させる。LegacyAPIを上手に活用して機能を実現する。
func (a Adapter) Receive(r *http.Request) (string, error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()
	return a.Legacy.Receive(r.Method, r.URL.Path, r.Header, b)
}
