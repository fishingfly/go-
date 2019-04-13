// 50_切片与底层数组的关系
package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := a[2:5:7]
	s1[2] = 900
	fmt.Println(s1)
	fmt.Println(a)
	s2 := s1[2:3]
	fmt.Println(s2)
}
