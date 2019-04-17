// 59_结构体普通变量初始化
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

func main() {
	// 顺序初始化，每个成员必须初始化
	var s1 Student = Student{1, "asd", 's', 89, "dfsfhsdj"}
	fmt.Println(s1)

	//指定成员初始化，未指定的自动赋值为0
	var s2 Student = Student{name: "test"}
	fmt.Println(s2)
}
