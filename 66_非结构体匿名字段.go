// 66_非结构体匿名字段
package main

import (
	"fmt"
)

type mystr string //自定义类型，给一个类型改名

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person //只有类型，没有名字，匿名字段，继承了Person的成员
	id     int
	add    string
	name   string
	int    // 基础类型字段的匿名字段
	mystr
}

func main() {
	var s Student
	s.int = 213
	s.mystr = "asdasd"
	fmt.Println(s)
}
