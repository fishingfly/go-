// 83_gobase3
package main

import (
	"fmt"
	"time"
)

func main() {
	name := "Eric"
	go func() {
		fmt.Printf("Hello, %s\n", name)
	}()
	name = "Harry"
	time.Sleep(time.Millisecond)

	//执行结果多为 hello harry。在赋值语句name = "Harry"执行之后，它上面的go函数才得以执行
	//可以选择吧最后两句换一下顺序看看执行结果 会是你想要的结果，但是仍然是不确定的，因为调度器的实时调度，我们是无法控制的
}
