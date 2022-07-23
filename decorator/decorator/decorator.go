package decorator

import (
	"encoding/base64"
	"strings"
)

// StringDecorator は文字列を装飾する
type StringDecorator interface {
	Decorate() string
}

type baseDecorator struct {
	s string
}

// NewBaseDecorator はオリジナルの文字列を返すデコレータを生成する
func NewBaseDecorator(s string) StringDecorator {
	return &baseDecorator{s: s}
}

func (bd *baseDecorator) Decorate() string {
	return bd.s
}

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

type reverseDecorator struct {
	d StringDecorator
}

// NewReverseDecorator は対象文字列を逆順にするデコレータを生成する
func NewReverseDecorator(d StringDecorator) StringDecorator {
	return &reverseDecorator{d: d}
}

func (rd *reverseDecorator) Decorate() string {
	var origin = []rune(rd.d.Decorate())
	var reversed strings.Builder
	for i := len(origin) - 1; i >= 0; i-- {
		reversed.WriteRune(origin[i])
	}
	return reversed.String()
}

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
