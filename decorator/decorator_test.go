// Package decorator_pattern
//
// [目的]
// Decoratorの目的は対象物を階層的にラップして機能拡張することです。
//
// [概要]
// ラップする際は 任意のDecoratorを組み合わせての拡張が可能で、この多様性がこのデザインの強みです。
//
// サンプル実装では次のDecoratorを用意し、任意の組み合わせで文字列を操作しています。
// - BaseDecorator: オリジナルの文字列を返すデコレータ
// - TrimDecorator: 対象文字列のスペースをトリムするデコレータ
// - ReverseDecorator: 対象文字列を逆順にするデコレータ
// - Base64Decorator: 対象文字列をBase64エンコードするデコレータ
//
// これらのデコレータを組み合わせることにより、多彩な文字操作を可能にしています。
package decorator_pattern

import (
	"github.com/koooyooo/go-design-pattern/decorator/decorator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecorate(t *testing.T) {
	// オリジナルの文字列に対し
	original := " Hello World! "

	// 様々なデコレータを組み合わせて、期待する結果を得る
	tests := []struct {
		name     string
		dec      decorator.StringDecorator
		expected string
	}{
		{
			name:     "original",
			dec:      decorator.NewBaseDecorator(original),
			expected: " Hello World! ",
		},
		{
			name:     "trim(original)",
			dec:      decorator.NewTrimDecorator(decorator.NewBaseDecorator(original)),
			expected: "Hello World!",
		},
		{
			name:     "reverse(trim(original))",
			dec:      decorator.NewReverseDecorator(decorator.NewTrimDecorator(decorator.NewBaseDecorator(original))),
			expected: "!dlroW olleH",
		},
		{
			name:     "base64(trim(original))",
			dec:      decorator.NewBase64Decorator(decorator.NewTrimDecorator(decorator.NewBaseDecorator(original))),
			expected: "SGVsbG8gV29ybGQh",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.dec.Decorate())
		})
	}
}
