package singleton

// 唯一のインスタンスを返却
func Instance() *singleton {
	return instance
}

// 唯一のインスタンス
var instance = &singleton{
	Value: "Hello",
}

// 構造体の宣言
type singleton struct {
	Value string
}
