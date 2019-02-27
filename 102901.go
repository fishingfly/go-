// 102901
package main

import (
	"fmt"
)

func test() {
	// 不同类型变量的声明
	var a int
	var b float32
	a, b = 10, 20
	fmt.Println(a, b)

	var (
		c int
		d float64
	)
	c, d = 19, 27.38
	fmt.Println(c, d)

	//	const i int = 10
	//	const j float32 = 34.5
	//	fmt.Println(i, j)

	const (
		i int     = 10 //这里面是可以自动推导类型的，可以省略int和float
		j float32 = 34.9
	)

	fmt.Println(i, j)
}
