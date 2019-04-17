// 67_结构体指针类型匿名字段
package main

import (
	"fmt"
)

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	*Person //指针类型
	id      int
	add     string
}

func main() {
	s1 := Student{&Person{"mike", 'm', 18}, 666, "bj"}
	fmt.Println("s1 = ", s1)
}
