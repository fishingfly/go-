// 71_指针变量的方法集
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
	// 结构变量是一个指针变量，他能够调用哪些方法，这些方法就是一个集合，简称方法集
	p := &Person{"mike", 'm', 19}

	// 内部做的转换，先把指针p,转成*p后再调用
	(*p).SetInfoValue("d", 'd', 89)
	p.SetInfoValue("d", 'd', 89)
}
