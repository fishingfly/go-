// 76_接口的继承
package main

import (
	"fmt"
)

type Humaner interface {
	sayhi()
}

type Personer interface { // 超集
	Humaner //匿名字段，继承了sayhi()
	sing(lrc string)
}

type Student struct {
	name string
	id   int
}

// student实现了sayhi()
func (tmp *Student) sayhi() {
	fmt.Println("tmp = ", *tmp)
}

func (tmp *Student) sing(lrc string) {
	fmt.Println("student sing: ", lrc)
}

func main() {
	// 定义一个接口类型的变量
	var i Personer
	s := &Student{"mike", 666}
	i = s
	i.sayhi()
	i.sing("sfddfd")
}
