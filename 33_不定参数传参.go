// 33_不定参数传参
package main

import (
	"fmt"
)

func MyFunc7(args ...int) {
	for _, data := range args {
		fmt.Println(data)
	}
}

func MyFunc6(args ...int) {
	MyFunc7(args...) // 将参数全部传过去
}

func MyFunc5(args ...int) {

}

func main() {
	MyFunc7(1, 23, 45, 6, 677)
}
