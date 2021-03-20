package immutable_fragile

// score はImmutableな得点(に見えて問題がある実装)
// 各所に脆弱性を埋め込んでいる
type score struct {
	// 状態がポインタなので、参照を持つものは書き換えが可能
	v *int
}

// 設定時に渡したポインタの操作で、呼び出し側が後から値を書き換えることが可能
func NewScore(v *int) *score {
	return &score{v}
}

// 取得時に獲得したポインタの操作で、呼び出し側が後から値を書き換えることが可能
func (s *score) Value() *int {
	return s.v
}
