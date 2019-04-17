// 70_值语义和引用语义
package main

import (
	"fmt"
)

type Person struct {
	name string
	sex  byte
	age  int
}

// 修改成员变量的值

//接收者为普通变量，非指针，值语义，一份拷贝
func (p Person) SetInfoValue(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Printf("setInfoValue &p = %p\n", p)
}

// 接受者为指针变量，引用传递
func (p *Person) SetInfoPointer(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Printf("setInfoPointer &p = %p\n", p)
}

func main() {
	s1 := Person{"sad", 's', 10}
	fmt.Printf("&s1 = %p\n", &s1)
	//值语义
	//	s1.SetInfoValue("mike", 'm', 18)
	//	fmt.Println("s1 = ", s1)

	//引用语义
	(&s1).SetInfoPointer("mike", 'm', 18)
	fmt.Println("s1 = ", s1) //打印内容
}
