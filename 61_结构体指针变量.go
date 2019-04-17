// 61_结构体指针变量
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
	// 指针有合法指向后才操作成员
	// 先定义一个普通的结构体变量
	var s Student
	var p *Student
	p = &s
	// 通过指针操作成员，p.id和(*p).id都行
	p.id = 91
	p.addr = "dhasd"
	p.age = 90
	fmt.Println("p = ", *p)

	// 通过new申请一个结构体
	p2 := new(Student)
	p2.addr = "adadas"
	fmt.Println("p2 = ", p2)
}
