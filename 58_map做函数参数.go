// 58_map做函数参数
package main

import (
	"fmt"
)

func test(m map[int]string) {
	delete(m, 1)
}

func main() {
	m := map[int]string{1: "21312", 2: "dsadasdas"}
	fmt.Println(m)
	test(m)
	fmt.Println(m)
}
