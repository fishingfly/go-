// 29_goto的使用
package main

import (
	"fmt"
)

func main() {
	// goto可以在任何地方使用，但是不能跨函数使用
	fmt.Println("1111111111")
	goto END // goto是关键字，End是用户起的名字，他叫标签
	fmt.Println("2122222222")
END:
	fmt.Println("33333333333")

}
