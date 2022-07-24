package singleton

import "sync"

// 構造体の宣言
type singleton2 struct {
	Value string
}

var instance2 *singleton2
var once sync.Once

// InstanceByOnce はインスタンスを必要に応じて生成して返却
func InstanceByOnce() *singleton2 {
	once.Do(func() {
		instance2 = &singleton2{
			Value: "Hello",
		}
	})
	return instance2
}
