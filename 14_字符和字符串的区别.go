// 14_字符和字符串的区别
package main

import (
	"fmt"
)

func main() {
	//	var ch byte
	var str string
	// 字符单引号，字符串双引号
	// 字符往往都只有一个字符，转义字符除外
	// 字符串是有1个或多个字符组成
	// 字符串是隐藏了一个结束符
	//	ch = 'a'
	str = "a" //有'a'和'\0'组成了一个字符串
	str = "hello go"
	fmt.Printf("str[0] = %c", str[0])
}
