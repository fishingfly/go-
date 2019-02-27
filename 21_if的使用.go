// 21_if的使用
package main

import (
	"fmt"
)

func main() {
	s := "王思聪"
	if s == "王思聪" { // 左括号和if同一行，且不可省略
		fmt.Println("Hello World!")
	}

	// if支持一个初始化语句,初始化语句和判断条件以分号分隔
	if a := 10; a == 10 {
		fmt.Println("a == 10")
	}
}
