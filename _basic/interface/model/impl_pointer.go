package model

//var _ Interface = InterfaceImpl{} // レシーバーがポインタなので実体では Interfaceを満たせていない
var _ Interface = &InterfaceImpl{}
var _ Interface = (*InterfaceImpl)(nil) // メモリも節約したい場合

// InterfaceImpl はそれを満足させる構造体
// Do関数のレシーバーがポインタ
type InterfaceImpl struct{}

func (i *InterfaceImpl) Do() string {
	return "do something"
}
