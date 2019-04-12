// 37_匿名函数和闭包
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	a := 10
	str := "mike"

	// 匿名函数，没有函数名字
	f1 := func() {
		fmt.Println("a = ", a)
		fmt.Println("str = ", str)
	}

	f1()

	// 给一个函数类型起别名
	type FuncType func() //函数没有参数，没有返回值
	//声明变量
	var f2 FuncType
	f2 = f1
	f2()

	// 定义匿名函数，同时调用
	func() {
		fmt.Println("a = ", a)
	}() // 后面的（）代表调用此匿名函数

	// 带参数的匿名函数
	f3 := func(i, j int) {
		fmt.Println("i=", i)
	}
	f3(1, 2)

	// 带参数的匿名函数,同时调用
	func(i, j int) {
		fmt.Println("i=", i)
	}(3, 4)

	// 匿名函数，有参有返回值
	func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(10, 20)
}
