package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if nums == nil || len(nums) == 0 {
		return []int{-1,-1}
	}
	begin := 0
	end := len(nums) - 1
	for {
		if nums[begin] < target {
			begin++
		}
		if nums[end] > target {
			end--
		}
		//什么时候跳出循环
		if end <= begin || (nums[begin] == target && nums[end] ==target) {
			break
		}
	}
	if begin == end {
		if nums[begin] != target {
			return []int{-1,-1}
		}
	}
	if begin > end {
		return []int{-1,-1}
	}
	return []int{begin,end}
}

func main() {
	fmt.Println(searchRange([]int{0,0,1,2,2,2,3}, 2))
	fmt.Println(searchRange([]int{2,2,2}, 2))
	fmt.Println(searchRange([]int{0,0}, 2))
	fmt.Println(searchRange([]int{}, 2))
	fmt.Println(searchRange([]int{1}, 2))
	fmt.Println(searchRange([]int{0,0,1,2,2,2,3}, 1))
	fmt.Println(searchRange([]int{0,0,1,2,2,2,3}, 0))
	fmt.Println(searchRange([]int{0,0,2,2,2,3}, 1))
}
