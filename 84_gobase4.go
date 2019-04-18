// 84_gobase4
package main

import (
	"fmt"
	"time"
)

//func main() {
//	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
//	for _, name := range names {
//		go func() {
//			fmt.Printf("Hello, %s\n", name)
//		}()
//	}
//	time.Sleep(time.Millisecond)
//	//原来是想对每个人打个hello的，却出现对同一个人五次hello
//}

func main() {
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		go func(who string) {
			fmt.Printf("Hello, %s\n", who)
		}(name)
	}
	time.Sleep(time.Millisecond)
	//利用传参数来解决上面问题，值类型参数传递进去就是传了一份拷贝，而要是传的引用类型就会出错
}
