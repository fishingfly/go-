// 46_数组值传递和地址传递
package main

import (
	"fmt"
)

func modify1(a [4]int) {
	a[0] = 0
}

func modify2(a *[4]int) {
	(*a)[0] = 0
}

func main() {
	fmt.Println("Hello World!")
	a := [4]int{1, 2, 3, 4}
	modify1(a)
	fmt.Println(a)
	modify2(&a)
	fmt.Println(a)
}
