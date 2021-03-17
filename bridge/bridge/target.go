package bridge

import (
	"encoding/json"
	"encoding/xml"

	"gopkg.in/yaml.v3"
)

var FormatJSON = &writingFormatJSON{}
var FormatYAML = &writingFormatYAML{}
var FormatXML = &writingFormatXML{}

type (
	WritingFormat interface {
		Marshal(interface{}) ([]byte, error)
	}

	writingFormatJSON struct{}
	writingFormatYAML struct{}
	writingFormatXML  struct{}
)

func (wj writingFormatJSON) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (wy writingFormatYAML) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (wx writingFormatXML) Marshal(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}
