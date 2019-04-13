// 56_map的遍历
package main

import (
	"fmt"
)

func main() {
	m := map[int]string{1: "222", 2: "dadad", 3: "dajfdf"}
	// 第一个返回值是key，第二个返回值是value，无序
	for key, value := range m {
		fmt.Printf("key = %d, value = %s", key, value)
	}

	// 如何判断一个key是否存在
	// 第一个返回值为key所对应的value，第二个是返回key是否存在，存在即为true
	value, ok := m[1]
	if ok == true {
		fmt.Print(value)
	} else {
		fmt.Println("没有这个key")
	}
}
