// 15_复数类型
package main

import (
	"fmt"
)

func main() {
	var t complex128
	t = 2.1 + 3.14i
	fmt.Println(t)

	// 自动推导类型
	t2 := 2.1 + 3.14i
	fmt.Printf("t2 type is %T", t2)

	// 通过内建函数，取实部和虚部
	fmt.Println("real(t2)=", real(t2), ", imag(t2) = ", imag(t2))
}
