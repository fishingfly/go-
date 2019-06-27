package main

import "fmt"

//把需要的数据放在前面即可，覆盖掉不需要的数据
func removeElement(nums []int, val int) int {
	indexNew := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[indexNew] = nums[i]
			indexNew++
		}
	}
	return indexNew
}

func main() {
	var arr = []int{0,1,2,2,3,0,4,2}
	len := removeElement(arr, 2)
	for i := 0; i < len; i++ {
		fmt.Println(arr[i])
	}
}