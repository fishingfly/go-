// 53_copy的使用
package main

import (
	"fmt"
)

func main() {
	src := []int{1, 2}
	dest := []int{3, 3, 3, 3, 3, 3, 33, 6}
	copy(dest, src)
	fmt.Print(src)
	fmt.Print(dest)
}
