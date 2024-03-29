package singleton

// 構造体の宣言
type singleton struct {
	Value string
}

// 唯一のインスタンスを用意
var instance = &singleton{
	Value: "Hello",
}

// Singletonの取得
func Instance() *singleton {
	// 唯一のインスタンスの参照を返却
	return instance
}
