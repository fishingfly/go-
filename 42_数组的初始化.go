// 42_数组的初始化
package main

import (
	"fmt"
)

func main() {
	// 声明定义同时赋值
	var a [5]int = [5]int{1, 2, 3, 4, 5} //全部初始化
	fmt.Print(a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Print(b)
	// 部分初始化,没有初始化的元素，自动赋值为0
	c := [5]int{1, 2, 3}
	fmt.Print(c)

	// 指定某个元素初始化
	d := [5]int{2: 10, 4: 20}
	fmt.Print(d)
}
