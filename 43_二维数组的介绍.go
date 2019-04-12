// 二维数组的介绍
package main

import (
	"fmt"
)

func main() {
	// 有多少个[]就是多少维
	var a [3][4]int
	k := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			k++
			a[i][j] = k
			fmt.Printf("a[%d][%d] = %d, ", i, j, a[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Println("a=", a)

	// 有三个元素，每个元素又是一堆数组
	b := [3][4]int{1: {1, 2, 3, 4}}
	fmt.Print(b)
}
