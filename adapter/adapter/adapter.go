package adapter

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	// レガシーAPI
	// (httpパッケージを活用せずに実装)
	LegacyAPI interface {
		Receive(method, url string, header map[string][]string, body []byte) (string, error)
	}
	// モダンAPI
	// (httpパッケージを活用して実装)
	ModernAPI interface {
		Receive(r *http.Request) (string, error)
	}
)

// レガシーAPI実装
type LegacyAPIImpl struct{}

// Receive はレガシーAPIの実装です
func (l LegacyAPIImpl) Receive(method, url string, header map[string][]string, body []byte) (string, error) {
	return fmt.Sprintf("legacy: %s %s", method, url), nil
}

// Adapter はアダプタの実装です。
type Adapter struct {
	// 委譲対象のLegacyAPIを保持
	Legacy LegacyAPI
}

// Receive は変換処理の実装でありモダンAPIのインターフェイスも満足させます。 LegacyAPIを上手に活用して機能を実現します。
func (a Adapter) Receive(r *http.Request) (string, error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()
	return a.Legacy.Receive(r.Method, r.URL.Path, r.Header, b)
}
