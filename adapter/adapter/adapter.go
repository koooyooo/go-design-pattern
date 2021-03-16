package adapter

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	// レガシーAPI
	LegacyAPI interface {
		Receive(method, url string, header map[string][]string, body []byte) (string, error)
	}
	// モダンAPI
	ModernAPI interface {
		Receive(r *http.Request) (string, error)
	}
)

// レガシーAPI実装
func NewLegacyAPI() LegacyAPI {
	return &legacyAPIImpl{}
}

type legacyAPIImpl struct{}

func (l legacyAPIImpl) Receive(method, url string, header map[string][]string, body []byte) (string, error) {
	return fmt.Sprintf("legacy: %s %s", method, url), nil
}

// Adapter定義
type Adapter struct {
	Legacy LegacyAPI
}

// 変換処理の実装
func (a Adapter) Receive(r *http.Request) (string, error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return a.Legacy.Receive(r.Method, r.URL.Path, r.Header, b)
}
