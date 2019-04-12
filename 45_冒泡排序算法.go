// 45_冒泡排序算法
package main

import (
	"fmt"
)

func BubbleSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func main() {
	a := []int{3, 4, 1, -1, 4342523, 4544543, 4, 324, 32432, 43245}
	BubbleSort(a)
	fmt.Print(a)
}
