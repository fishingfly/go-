// 65_同名字段
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
	name   string
}

func main() {

	var s Student

	// 默认规则:如果能在本作用域找到此成员，就操作此成员，如果没有找到，找到继承的字段
	s.name = "aaaa" // 操作的是Student的字段不是person的字段
	// 显示调用
	s.Person.name = "bbbbb"
	fmt.Println(s)

}
