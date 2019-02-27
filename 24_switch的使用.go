// 24_switch的使用
package main

import (
	"fmt"
)

func main() {
	num := 18
	switch num { // switch后面写的是变量本身
	case 1:
		fmt.Printf("按下的是%d\n", num)
		break // go语言保留了break关键字，跳出当前循环，但是默认不写就包含了，在switch中
		//		fallthrough // 不跳出switch语句，后面无条件执行
	case 2:
		fmt.Printf("按下的w是%d\n", num)
		break
	case 3:
		fmt.Printf("按下的是%d\n", num)
		break
	case 4:
		fmt.Printf("按下的是%d\n", num)
		break
	default:
		fmt.Printf("按下的是%xxx")
	}
}
