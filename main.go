package main

import "fmt"

func main() {
	// プリミティブのキャスト
	var a = (int32)(int64(32))
	fmt.Println(a)

	// インターフェイスのキャスト
	var in interface{} = "hello"
	var b, ok = in.(string)
	if ok {
		fmt.Println(b)
	}
}
