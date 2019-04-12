// 35_函数类型
package main

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

//函数也是一种数据类型，通过type给一个函数类型起名
// FuncType它是一个函数类型
type FuncType func(int, int) int

func main() {
	var result int
	result = Add(1, 1)
	fmt.Println("result = ", result)

	// 声明一个函数类型的变量，变量名叫fTest
	var fTest FuncType
	fTest = Add
	result = fTest(10, 20)
	fmt.Println("result 2 = ", result)

	fTest = minus
	result = fTest(10, 5)
	fmt.Println("result 3 = ", result)
}
