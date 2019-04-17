// 63_匿名字段的初始化
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
	Person //只有类型，没有名字，匿名字段，继承了Person的成员
	id     int
	add    string
}

func main() {
	// 顺序初始化
	var s1 Student = Student{Person{"mike", 'm', 19}, 1, "dssa"}
	fmt.Println("s1 == ", s1)

	// 自动推导类型
	s2 := Student{Person{"mike", 'm', 19}, 1, "dssa"}
	fmt.Println("s2 == ", s2)
}
