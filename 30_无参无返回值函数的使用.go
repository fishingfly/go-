// 30_无参无返回值函数的使用
package main

import (
	"fmt"
)

//无参无返回值函数的定义
func MyFunc() {
	a := 666
	fmt.Println(a)
}

func main() {
	MyFunc()
}
