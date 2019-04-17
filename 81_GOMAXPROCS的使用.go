// 81_GOMAXPROCS的使用
package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(1) //指定以单核运算
	fmt.Println("n = ", n)

	for {
		go fmt.Print(1)
		fmt.Print(0)
	}

}
