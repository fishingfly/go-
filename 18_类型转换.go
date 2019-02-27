// 18_类型转换
package main

import (
	"fmt"
)

func main() {
	var flag bool
	flag = true
	fmt.Printf("flag = %d\n", flag)
	// bool类型不能转换为整型
	// 0就是假，1就是真
	// 整型也不能转换为bool
	// bool是不兼容类型

	var ch byte
	ch = 'a' //字符型本质上就是整形
	var t int
	t = int(ch)
	fmt.Println("t = ", t)
}
