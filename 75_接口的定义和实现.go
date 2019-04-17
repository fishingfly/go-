// 75_接口的定义和实现
package main

import (
	"fmt"
)

//定义接口类型
type Humaner interface {
	//方法，只有声明，没有实现，由别的类型（自定义类型）实现
	sayhi()
}

type Student struct {
	name string
	id   int
}

// student实现了此方法
func (tmp *Student) sayhi() {
	fmt.Printf("Student[]\n")
}

type Teacher struct {
	addr  string
	group string
}

//Teacher实现了此方法
func (tmp *Teacher) sayhi() {
	fmt.Printf("Teacher[%s, %s] say hi\n", tmp.addr, tmp.group)
}

type MyStr string

//MyStr实现此方法
func (tmp *MyStr) sayhi() {
	fmt.Printf("MyStr[%s] sayhi\n", *tmp)
}

func main() {
	// 定义接口类型的变量
	var i Humaner

	//只是实现了此接口方法的类型，那么这个类型的变量（接收者类型）就可以给i赋值
	s := &Student{"mike", 666}
	i = s
	s.sayhi()

	t := &Teacher{"bj", "dsdsd"}
	i = t
	i.sayhi()
}
