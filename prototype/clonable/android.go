package clonable

// Android はClone可能な構造体
type Android struct {
	Name string
	Log  []string
}

// 自身の複製を生成するメソッド
func (a *Android) Clone() Android {
	// Goでは値(非ポインタ)の代入だけでクローン可能
	b := *a
	return b
}
