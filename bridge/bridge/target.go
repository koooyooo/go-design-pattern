package bridge

import (
	"fmt"
	"os"
)

type (
	WritingTarget struct {
		Write  func(b []byte) (string, error)
		Format WritingFormat
	}
)

func WriteStdout(b []byte) (string, error) {
	s := fmt.Sprintf("[stdout] %s", string(b))
	_, err := fmt.Fprint(os.Stdout, s)
	return s, err
}

func WriteStderr(b []byte) (string, error) {
	s := fmt.Sprintf("[stderr] %s", string(b))
	_, err := fmt.Fprint(os.Stderr, s)
	return s, err
}

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
