// 51_append函数的使用
package main

import (
	"fmt"
)

func main() {
	s1 := []int{0, 1, 2}
	s1 = append(s1, 3)
	s1 = append(s1, 3)
	s1 = append(s1, 3)
	fmt.Print(s1)
	fmt.Printf("len(s1)=%d, cap(s1)=%d", len(s1), cap(s1))
}
