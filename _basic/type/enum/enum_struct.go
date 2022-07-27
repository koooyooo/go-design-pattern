package enum

import "fmt"

// 単純型ではないので constでは定義できない
var (
	January  = &structEnum{no: 1, name: "January"}
	February = &structEnum{no: 2, name: "February"}
	March    = &structEnum{no: 3, name: "March"}
)

// 構造体なので任意の情報を管理できる
type structEnum struct {
	no   int
	name string
}

// No では値を書き駆られないよう フィールドはExportせずに、Getterを実装する
func (s *structEnum) No() int {
	return s.no
}

// String は適切な文字列表現を行う。将来的な変更を想定しここでの文字列表現は画面等のViewでは用いず主にデバッグ等で用いる
func (s *structEnum) String() string {
	return fmt.Sprintf("[%d] %s", s.no, s.name)
}
