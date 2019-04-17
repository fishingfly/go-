// 62_结构体的比较和赋值
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
	p1 := Student{1, "stud", 'm', 18, "dsadasd"}
	p2 := Student{1, "stud", 'm', 18, "dsadasd"}
	p3 := Student{2, "stud", 'm', 18, "dsadasd"}
	fmt.Println("p1 == p2 : ", p1 == p2)
	fmt.Println("p3 == p2 : ", p3 == p2)

	// 同类型的2个结构体变量可以相互赋值
	var tmp Student
	tmp = p1
	fmt.Print(tmp)
}
