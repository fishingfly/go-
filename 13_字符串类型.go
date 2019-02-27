// 13_字符串类型
package main

import (
	"fmt"
)

func main() {
	var str1 string
	str1 = "aaaaaa"
	fmt.Println(str1)

	// 自动推到类型
	str2 := "mike"
	fmt.Printf("str2类型是%T\n", str2)
	// 内建函数，len()可以测字符串的长度
	fmt.Println("len(str2) = ", len(str2))
}
