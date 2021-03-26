package composite

import (
	"fmt"
	"log"
	"testing"
)

func TestDo(t *testing.T) {
	tf, err := NewTreeFile("../../README.md")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tf.fileInfo.Mode().String())
	fmt.Println(tf.fileInfo.Sys())
	fmt.Println(tf.fileInfo.Name())
}
