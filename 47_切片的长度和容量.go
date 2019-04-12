// 47_切片的长度和容量
package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 0, 0}
	s := a[0:3:5]
	fmt.Println("s = ", s)
	fmt.Println("len(s) = ", len(s)) //长度
	fmt.Println("cap(s) = ", cap(s)) //容量
}
