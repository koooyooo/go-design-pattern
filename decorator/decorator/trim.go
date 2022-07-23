package decorator

import "strings"

type trimDecorator struct {
	d StringDecorator
}

// NewTrimDecorator は対象文字列のスペースをトリムするデコレータを生成する
func NewTrimDecorator(d StringDecorator) StringDecorator {
	return &trimDecorator{d: d}
}

func (td *trimDecorator) Decorate() string {
	return strings.TrimSpace(td.d.Decorate())
}
