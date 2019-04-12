// 40_defer的使用
package main

import (
	"fmt"
)

func main() {
	//延迟执行,main函数结束前调用
	defer fmt.Println("ssssssssss")
	fmt.Println("Hello World!")

}
