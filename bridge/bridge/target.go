package bridge

import (
	"fmt"
	"os"
)

type (
	// WritingTarget は出力対象を表現します
	WritingTarget struct {
		// 抽象メソッド相当の関数
		Write
		// 委譲先のフォーマット
		Format WritingFormat
	}
	// 出力関数
	Write func(b []byte) (string, error)
)

// WriteOut は自身の抽象メソッド(関数)と委譲先のフォーマットを活用して出力を行います
func (wb WritingTarget) WriteOut(v interface{}) (string, error) {
	b, err := wb.Format.Marshal(v)
	if err != nil {
		return "", err
	}
	s, err := wb.Write(b)
	if err != nil {
		return "", err
	}
	return s, nil
}

// 出力関数の標準出力実装
var WriteStdout Write = func(b []byte) (string, error) {
	s := fmt.Sprintf("[stdout] %s", string(b))
	_, err := fmt.Fprint(os.Stdout, s)
	return s, err
}

// 出力関数の標準エラー出力実装
var WriteStderr Write = func(b []byte) (string, error) {
	s := fmt.Sprintf("[stderr] %s", string(b))
	_, err := fmt.Fprint(os.Stderr, s)
	return s, err
}
