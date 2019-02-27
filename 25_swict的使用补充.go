// 25_swict的使用补充
package main

import (
	"fmt"
)

func main() {
	//支持一个初始化语句，初始化语句和变量本身，以分号分隔
	switch num := 7; num { // switch后面写的是变量本身
	case 1:
		fmt.Printf("按下的是%d\n", num)
		break // go语言保留了break关键字，跳出当前循环，但是默认不写就包含了，在switch中
		//		fallthrough // 不跳出switch语句，后面无条件执行
	case 2:
		fmt.Printf("按下的w是%d\n", num)
		break
	case 3, 7, 8:
		fmt.Printf("按下的是%d\n", num)
		break
	case 4:
		fmt.Printf("按下的是%d\niu fgvxs", num)
		break
	default:
		fmt.Printf("按下的是%xxx")
	}

	// 区别其他语言的新用法
	score := 85
	switch { // 可以没有条件
	case score > 90: // case
		fmt.Println("优秀")
	case score > 80:
		fmt.Println("良好")
	case score > 70:
		fmt.Println("一般")
	case score > 60:
		fmt.Println("及格")
	default:
		fmt.Println("差")
	}

}
