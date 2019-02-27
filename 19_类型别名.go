// 19_类型别名
package main

import (
	"fmt"
)

func main() {
	// 给int64起一个别名叫bigint
	type bigint int64
	var a bigint //等价于int64 a

	fmt.Printf("%T", a)

}
