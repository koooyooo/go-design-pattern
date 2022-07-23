package bridge

import (
	"fmt"
	"os"
)

// Destination は出力先を表す関数
type Destination func(b []byte) (string, error)

// DestStdout は出力関数の標準出力実装
var DestStdout Destination = func(b []byte) (string, error) {
	s := fmt.Sprintf("[stdout] %s", string(b))
	_, err := fmt.Fprint(os.Stdout, s)
	return s, err
}

// DestStderr は出力関数の標準エラー出力実装
var DestStderr Destination = func(b []byte) (string, error) {
	s := fmt.Sprintf("[stderr] %s", string(b))
	_, err := fmt.Fprint(os.Stderr, s)
	return s, err
}
