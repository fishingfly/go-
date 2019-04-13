// 57_map的删除
package main

import (
	"fmt"
)

func main() {
	m := map[int]string{1: "sdsd", 3: "sadsad"}
	fmt.Print(m)

	delete(m, 1) // 删除key为1的键值对
	fmt.Print(m)
}
