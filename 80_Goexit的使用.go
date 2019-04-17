// 80_Goexit的使用
package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("ccccc")

	// return//终止此函数
	runtime.Goexit() //终止所在的协程
	fmt.Println("ddddddd")
}

func main() {
	//创建新建的协程
	go func() {
		fmt.Println("aaaaaa")
		test()
		fmt.Println("bbbbbbb")
	}()

	for {

	}

	fmt.Println("Hello World!")
}
