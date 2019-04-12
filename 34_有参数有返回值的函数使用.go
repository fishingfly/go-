// 34_有参数有返回值的函数使用
package main

import (
	"fmt"
)

func MyFunc08(a, b int) (max, min int) {
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	return //有返回值的地方一定要有return
}

func main() {
	a, b := MyFunc08(5, 6)
	fmt.Println(a, b)
}
