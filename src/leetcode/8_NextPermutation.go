package main

import "fmt"

func nextPermutation(nums []int)  {
	if nums == nil || len(nums) == 0 {
		return
	}
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i] > nums[i - 1] {
			nums[i],nums[i - 1] = nums[i - 1],nums[i]
			return
		}
	}
	// 逆转所有的数字
	for i := 0; i < len(nums)/2; i++ {
		nums[i],nums[len(nums) - 1 - i] = nums[len(nums) - 1 - i],nums[i]
	}
}



func main() {
	nums := []int{1,1,5}
	nextPermutation(nums)
	fmt.Println(nums)
}
