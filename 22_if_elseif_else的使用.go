// 22_if_elseif_else的使用
package main

import (
	"fmt"
)

func main() {
	//1
	a := 13
	//	if a == 10 {
	//		fmt.Println("a=10")
	//	} else {
	//		fmt.Println("Hello World!")
	//	}

	if a == 10 {
		fmt.Println("a=10")
	} else if a < 10 {
		fmt.Println("a < 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else {
		fmt.Println("不可能")
	}

}
