// 64_成员操作
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
	s1 := Student{Person{"saddas", 's', 78}, 89, "sads"}
	s1.age = 79889
	fmt.Println(s1)
	s1.Person = Person{"ssss", 'a', 8989}
	fmt.Println(s1)
}
