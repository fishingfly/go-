// 49_切片的截取
package main

import (
	"fmt"
)

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//[low:high:max] 取下标low开始的元素，len = high - low;cap = max-len
	s1 := array[:] // [0:len(array):len(array)] 不指定容量，容量就和长度一样
	fmt.Print(s1)
	fmt.Printf("len(s1):%d,cab(s1):%d", len(s1), cap(s1))

	//操作某个元素和数组的方式一样

	s2 := array[3:6:6]
	fmt.Print(s2)
	fmt.Printf("len(s2):%d,cab(s2):%d", len(s2), cap(s2))

	s3 := array[:6] // 从0开始取6个元素，容量也是6
	fmt.Print(s3)
	fmt.Printf("len(s3):%d,cab(s3):%d", len(s3), cap(s3))
}
