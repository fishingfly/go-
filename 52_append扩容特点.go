// 52_append扩容特点
package main

import (
	"fmt"
)

func main() {
	array := []int{}
	oldCap := cap(array)
	for i := 0; i < 8; i++ {
		array = append(array, i)
		if newCap := cap(array); oldCap < newCap {
			fmt.Printf("oldCap=%d,newCap=%d\n", oldCap, newCap)
			oldCap = newCap

		}
	}
}
