// 54_切片做函数参数
package main

import (
	"fmt"
)

func initial(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] = i
	}
}

func main() {
	n := 10
	s := make([]int, n)
	initial(s)
	fmt.Print(s)

}
