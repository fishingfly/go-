package main

import "fmt"

func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	return solve(nums, 0,len(nums) - 1,target)
}
func solve(nums []int, begin,end,target int) int {
	for ; begin <= end; {
		if nums[begin] < nums[end] {//这段是升序的，直接用二分查找
			return TwoPintLookup(nums,begin,end,target)
		} else if nums[begin] > nums[end] {//这段存在反转数组，需要另外考虑，最终都会用二分查找
			medium := (end - begin)/2 + begin
			resultHead := solve(nums, begin,medium,target)
			resultTail := solve(nums, medium + 1,end,target)
			if resultHead != -1 {
				return resultHead
			}
			if resultTail != -1 {
				return resultTail
			}
			break
		} else {
			if nums[begin] == target{
				return begin
			} else {
				break
			}
		}
	}
	return -1
}



//数组是递增的
func TwoPintLookup(nums []int, begin,end,target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	for {
		medium := (end - begin)/2 + begin
		if nums[medium] == target {
			return medium
		}
		if begin >= end {
			return -1
		}
		if nums[begin] > target || nums[end] < target {
			return -1
		} else {
			if nums[medium] > target {
				end = medium
			} else {
				begin =medium + 1
			}
		}
	}
}

func main() {
	nums := []int{3,4,5,6,7,0,1,2}
	fmt.Println(search(nums, 3))
	fmt.Println(search(nums, 2))
	fmt.Println(search(nums, 0))
	fmt.Println(search(nums, 7))
	nums = nums[:1]
	fmt.Println(search(nums, 7))
	fmt.Println(search(nums, 3))
	fmt.Println(search(nums, 4))
}