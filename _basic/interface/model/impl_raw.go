package model

var _ Interface = InterfaceImplRaw{}
var _ Interface = &InterfaceImplRaw{}
var _ Interface = (*InterfaceImplRaw)(nil) // メモリも節約したい場合

// InterfaceImplRaw はそれを満足させる構造体
// Do関数のレシーバーが値
type InterfaceImplRaw struct{}

func (i InterfaceImplRaw) Do() string {
	return "do something"
}
