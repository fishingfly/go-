// 48_切片的创建
package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 5, 10)
	fmt.Printf("len(s1):%d,cab(s1):%d", len(s1), cap(s1))
	s2 := make([]int, 6, 9)
	fmt.Printf("len(s2):%d,cab(s2):%d", len(s2), cap(s2))
}
