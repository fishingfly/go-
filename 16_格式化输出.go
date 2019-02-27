// 16_格式化输出
package main

import (
	"fmt"
)

func main() {
	a := 10
	b := "adb"
	c := 'a'
	d := 3.14
	// %d 整形格式
	// %s 字符串格式
	// %c 字符个数
	// %f 浮点型个数
	fmt.Printf("%T,%T,%T,%T\n", a, b, c, d)
	fmt.Printf("%d,%s,%c,%f\n", a, b, c, d)
	fmt.Printf("%v,%v,%v,%v\n", a, b, c, d) // %v自动匹配格式输出，不是很智能
}
