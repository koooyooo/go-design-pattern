package decorator

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
