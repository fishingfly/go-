// 73_方法的重写
package main

import (
	"fmt"
)

type Person struct {
	name string //名字
	sex  string // 性别，字符串
	age  int    //年龄
}

// Person类型，实现一个方法
func (tmp *Person) PrintInfo() {
	fmt.Printf("name = %s, sex=%c, age=%d\n", tmp.name, tmp.sex, tmp.age)
}

// 有个学生，继承Person字段，成员和方法都继承了
type Student struct {
	Person //匿名字段
	id     int
	addr   string
}

//student也实现了一个方法，这个方法和Person方法同名，这种方法叫重写
func (tmp *Student) PrintInfo() {
	fmt.Println("tmp = ", tmp)
}

func main() {
	s := Student{Person{"mike", "m", 18}, 888, "bj"}
	//就近原则：先找到本作用域的方法，找不到再用继承的方法
	s.PrintInfo()

}
