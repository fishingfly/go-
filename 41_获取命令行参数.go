// 41_获取命令行参数
package main

import (
	"fmt"
	"os"
)

// 给包起别名
//import io "fmt"

// 忽略此包
//import _ "fmt"

func main() {
	// 接收用户传递的参数，都是以字符串方式传递的
	list := os.Args
	n := len(list)
	fmt.Println(n)
	for i := 0; i < n; i++ {
		fmt.Printf("list[%d] = %s\n", i, list[i])
	}
}
