package model

var _ Interface = InterfaceImplVal{}
var _ Interface = &InterfaceImplVal{}
var _ Interface = (*InterfaceImplVal)(nil) // メモリも節約したい場合

// InterfaceImplVal はそれを満足させる構造体
// Do関数のレシーバーが値
type InterfaceImplVal struct{}

func (i InterfaceImplVal) Do() string {
	return "do something"
}
