// 74_方法值
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
	fmt.Println("p =", *p)
}

func main() {
	p := Person{"mike", 'm', 19}
	p.SetInfoPointer("aaa", 'd', 109) //传统的调用方式

	//保存方式入口地址
	pFunc := p.SetInfoPointer // 这个就是方法值，调用函数时，无需再传递接收者，隐藏了接收者，等价于p.SetInfoPointer()
	pFunc("aadda", 'D', 1009)

}
