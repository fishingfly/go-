// 68_面向对象和面向过程的区别
package main

import (
	"fmt"
)

//面向过程
func Add(a, b int) int {
	return a + b
}

// 面向对象，方法：给某个类型绑定一个函数
type long int

func (tmp long) Add01(other long) long {
	return tmp + other
}

func main() {
	var result int
	result = Add(1, 1)
	fmt.Println("result = ", result)

	// 定义一个变量
	var a long = 2
	//调用方法格式： 变量名.函数(所需参数)
	r := a.Add01(3)
	fmt.Println("r = ", r)
}
