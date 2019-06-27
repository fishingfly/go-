package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}


func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	lastNum := nums[0]
	lenNew := 1
	for i := 1; i < len(nums); i++ {
		if lastNum == nums[i] {
			continue
		} else {
			lastNum = nums[i]
			nums[lenNew] = nums[i]
			lenNew++
		}
	}
	return lenNew
}

func main()  {
	arrays := []int{1,1,2}
	fmt.Println(removeDuplicates(arrays))
}


