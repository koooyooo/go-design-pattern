package enum

const (
	Hello enum = iota + 1 // 初期値 0と判別するため +1スタート
	World
)

// 直接生成をされないよう Exportしない
type enum int

// switchを利用して文字列を返す。
// 後の挿入で値がズレても問題ない様に switch条件を数値リテラルではなくenum値とする
func (e enum) String() string {
	switch e {
	case Hello:
		return "Hello"
	case World:
		return "World"
	default:
		return "Unknown"
	}
}
