package singleton

// 通常の構造体の取得
func NewNormal() *normal {
	// 毎回生成して返却
	return &normal{
		Value: "Hello",
	}
}

// 通常の構造体の宣言
type normal struct {
	Value string
}
