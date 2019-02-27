// 27_range的使用
package main

import (
	"fmt"
)

func main() {
	str := "abc"
	//	for i := 0; i < len(str); i++ {
	//		fmt.Printf("%c", str[i])
	//	}
	// 迭代打印每个元素，默认返回2个值：一个是元素的位置，一个是元素本身
	for i, data := range str {
		fmt.Printf("str[%d]=%c\n", i, data)
	}

	for i, _ := range str { // 第2个返回值默认丢弃，返回元素的位置（下标）
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}
}
