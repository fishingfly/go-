// 08_多个变量或常量定义
package main

import (
	"fmt"
)

func main() {
	//不同类型变量的声明
	//	var a int
	//	var b float64
	var (
		a int
		b float64
	)
	a, b = 10, 3.14
	fmt.Println(a, b)

	//	const i int = 10
	//	const j float64 = 3.14
	//	const (
	//		i int     = 10
	//		j float64 = 3.14
	//	)
	//可以自动推导类型
	const (
		i = 10
		j = 3.14
	)
	fmt.Println(i, j)
}
