package main

func canJump(nums []int) bool {
	return start(nums, 0)
}

func start(nums[]int, currentPos int) bool{
	if currentPos > len(nums) - 1 {
		return false
	}
	if currentPos == len(nums) - 1 {
		return true
	}
	if nums[currentPos] == 0 {
		return false
	}
	for i := nums[currentPos]; i > 0; i++ {
		if start(nums, currentPos + i) {
			return true
		}
	}
	return false
}

func main() {
}
