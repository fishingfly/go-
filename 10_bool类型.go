// 10_bool类型
package main

import (
	"fmt"
)

func main() {
	// 没有初始化时初始值为false
	var a bool
	a = true
	fmt.Println(a)

	// 自动推导类型
	var b = false
	fmt.Println(b)
	c := false
	fmt.Println(c)
}
