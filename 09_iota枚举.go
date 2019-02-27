// 09_iota枚举
package main

import (
	"fmt"
)

func main() {
	// iota常量自动生成器，每隔一行，自动累加1
	// iota给常量赋值
	// iota遇到const给赋值
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, b, c)
	// iota遇到const给赋值
	const d = iota
	fmt.Println(d)

	// 可以只写一个iota
	const (
		a1 = iota
		b1
		c1
	)
	fmt.Println(a1, b1, c1)
	// 如果同一行，值都一样
	const (
		i          = iota
		j1, j2, j3 = iota, iota, iota
		k          = iota
	)
	fmt.Println(i, j1, j2, j3, k)

}
