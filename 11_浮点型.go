// 11_浮点型
package main

import (
	"fmt"
)

func main() {
	var f1 float32
	f1 = 3.14
	fmt.Println(f1)
	// 自动推到类型, 默认是float64，float64比float32精确
	f2 := 3.14
	fmt.Println(f2)
}
