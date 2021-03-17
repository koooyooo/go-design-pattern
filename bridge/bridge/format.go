package bridge

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

var FormatJSON = &writingFormatJSON{}
var FormatYAML = &writingFormatYAML{}

type (
	WritingFormat interface {
		Marshal(interface{}) ([]byte, error)
	}

	writingFormatJSON struct{}
	writingFormatYAML struct{}
)

func (wj writingFormatJSON) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (wy writingFormatYAML) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}
