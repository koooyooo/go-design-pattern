package immutable

// score はImmutableな得点
// 可視性はpackage内に収めて外から使わせる
type score struct {
	// 値はImmutableな型を推奨 (不可能ならコンストラクタでの設定時とGetterでの返却時にコピーした値を用いて影響力を断ち切る)
	// 可視性もpackage内に留める
	v int
}

// コンストラクタで値を設定し Setterを設けない
func NewScore(v int) *score {
	return &score{v}
}

// 取得する値はポインタ渡しを控えるかコピーして渡す
func (s score) Value() int {
	return s.v
}
