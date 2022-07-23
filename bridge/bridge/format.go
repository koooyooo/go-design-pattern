package bridge

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

var (
	// FormatJSON はJSON形式で出力するフォーマット
	FormatJSON Format = &formatJSON{}

	// FormatYAML はYAML形式で出力するフォーマット
	FormatYAML Format = &formatYAML{}
)

type (
	// Format は入出力フォーマットを表すインターフェース
	Format interface {
		Marshal(interface{}) ([]byte, error)
		Unmarshal([]byte, interface{}) error
	}

	formatJSON struct{}
	formatYAML struct{}
)

func (wj formatJSON) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (wj formatJSON) Unmarshal(b []byte, v any) error {
	return json.Unmarshal(b, v)
}

func (wy formatYAML) Marshal(v any) ([]byte, error) {
	return yaml.Marshal(v)
}

func (wy formatYAML) Unmarshal(b []byte, v any) error {
	return yaml.Unmarshal(b, v)
}
