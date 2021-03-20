package immutable_fragile

type score struct {
	v *int
}

func NewScore(v int) *score {
	return &score{&v}
}

func (s *score) Value() *int {
	return s.v
}
