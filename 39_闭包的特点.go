// 39_闭包的特点
package main

import (
	"fmt"
)

// 函数的返回值是一个匿名函数返回一个函数类型
func test02() func int {
	var x int //没有初始化，值为0
	return func() int {
		x++
		return x * x
	}
}

func test01() int {
	// 函数被调用时，x才分配空间，才初始化为0
	var x int
	x++
	return x * x // 函数调用完毕，x自动释放
}

func main() {
//	fmt.Println(test01())
//	fmt.Println(test01())
//	fmt.Println(test01())
//	fmt.Println(test01())
// 返回值是一个匿名函数，返回一个函数类型，通过F来调用返回的匿名函数，f来调用闭包函数
//他不关心这些捕获了的变量是否已经超出了作用域
// 所以只有闭包还在使用它，这些变量就还会一直存在
	f := test02()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	
	

}
