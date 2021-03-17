package non_singleton

// 通常の構造体の取得
func New() *nonSingle {
	// 毎回生成して返却
	return &nonSingle{
		Value: "Hello",
	}
}

// 通常の構造体の宣言
type nonSingle struct {
	Value string
}
