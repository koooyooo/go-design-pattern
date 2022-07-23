package decorator

import "strings"

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
