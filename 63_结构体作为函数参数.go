// 63_结构体作为函数参数
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

func test01(s Student) {
	s.id = 999
	fmt.Println("test01: ", s)
}

func test02(s *Student) {
	s.id = 999
	fmt.Println("test02: ", *s)
}

func main() {
	s := Student{12, "dafd", 's', 23, "加油"}
	test01(s) // 值传递，形参无法改实参
	fmt.Println("s = ", s)
	test02(&s) // 地址传递，形参P3改实参
	fmt.Println("s = ", s)
}
