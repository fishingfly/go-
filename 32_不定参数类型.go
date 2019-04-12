// 32_不定参数类型
package main

import (
	"fmt"
)

// 不定参数
//注意不定参数，只能放在形参中的最后一个参数
func MyFunc01(args ...int) { //传递的实参是可以0或3个
	fmt.Println("len(args)=", len(args))
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}

	for _, data := range args {
		fmt.Println(data)
	}
}

func main() {
	MyFunc01(1, 2, 3, 4, 5, 6)
}
