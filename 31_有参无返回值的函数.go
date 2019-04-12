// 31_有参无返回值的函数
package main

import (
	"fmt"
)

// 有参数无返回值函数的定义,以及一些形参的常见写法

func MyFunc(a int) {
	fmt.Println(a)
}

func MyFunc1(a int, b int) {
	fmt.Println(a, b)
}

func MyFunc4(a, b int) {
	fmt.Println(a, b)
}

func MyFunc2(a int, b string, c float32) {
	fmt.Println(a, b, c)
}

func MyFunc3(a, b string, c float64, d, e int) {

}

func main() {
	MyFunc(9999)
}
