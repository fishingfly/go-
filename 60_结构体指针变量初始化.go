// 60_结构体指针变量初始化
package main

import (
	"fmt"
)

// 定义一个结构体类型
type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	var p1 *Student = &Student{1, "stud", 'm', 18, "dsadasd"}
	fmt.Println("p1 = ", *p1)

}
