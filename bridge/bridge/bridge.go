package bridge

type (
	// DestinationFormatBridge は出力対象を表す構造体
	DestinationFormatBridge struct {
		// 抽象先を表す関数（関数実装）
		Destination
		// フォーマットを表すインターフェイス（インターフェイス実装）
		Format
	}
)

// WriteOut は自身の抽象メソッド(関数)と委譲先のフォーマットを活用して出力を行う
func (dfb DestinationFormatBridge) WriteOut(v interface{}) (string, error) {
	b, err := dfb.Format.Marshal(v)
	if err != nil {
		return "", err
	}
	s, err := dfb.Destination(b)
	if err != nil {
		return "", err
	}
	return s, nil
}
