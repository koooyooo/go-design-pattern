package decorator

import "encoding/base64"

// NewBase64Decorator は対象文字列をBase64エンコードするデコレータを生成する
func NewBase64Decorator(d StringDecorator) StringDecorator {
	return &base64Decorator{d: d}
}

type base64Decorator struct {
	d StringDecorator
}

func (td *base64Decorator) Decorate() string {
	return base64.StdEncoding.EncodeToString([]byte(td.d.Decorate()))
}
