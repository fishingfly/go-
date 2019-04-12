// 44_随机数的使用
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置种子，只需要1次
	// 如果种子参数一样，每次运行程序产生的随机数都一样
	rand.Seed(time.Now().UnixNano())
	// 产生随机数
	for i := 0; i < 5; i++ {
		fmt.Print("rand = ", rand.Int())
		fmt.Print("rand = ", rand.Intn(100))
	}
}
