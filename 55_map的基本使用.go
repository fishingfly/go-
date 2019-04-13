// 55_map的基本使用
package main

import (
	"fmt"
)

func main() {
	mapTemp := map[int]string{1: "like", 2: "test"}
	fmt.Println(mapTemp)
	mapTemp[1] = "think"
	fmt.Println(mapTemp)

	var m1 map[int]string
	// 对于map只有len没有cap
	fmt.Println(len(m1))

	// 通过make可以创建map，指定长度，但是里面缺一个数据都没有
	m2 := make(map[int]string, 2)
	m2[1] = "test"
	m2[2] = "test"
	m2[3] = "test" //map底层自动扩容
	fmt.Println(m2)

}
