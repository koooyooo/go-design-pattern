package bridge

import (
	"fmt"
	"io"
	"os"
)

var TargetString = &writingTargetString{}
var TargetStdout = &writingTargetStdout{}

type (
	WritingTarget io.Writer

	writingTargetString struct {
		b []byte
	}
	writingTargetStdout struct {
	}
)

func (wts *writingTargetString) Write(p []byte) (int, error) {
	wts.b = append(wts.b, p...)
	return len(p), nil
}

func (wts writingTargetString) String() string {
	return string(wts.b)
}

func (wts *writingTargetStdout) Write(p []byte) (int, error) {
	n, err := fmt.Fprint(os.Stdout, string(p))
	if err != nil {
		return n, err
	}
	return len(p), nil
}
