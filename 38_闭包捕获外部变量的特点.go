// 38_闭包捕获外部变量的特点
package main

import (
	"fmt"
)

func main() {
	// 闭包以引用方式捕获外变量
	a := 10
	str := "mike"

	func() {
		a = 666
		str = "go"
		fmt.Println("a = ", a)
		fmt.Println("str = ", str)
	}()
	fmt.Println("a = ", a)
	fmt.Println("str = ", str)
}
