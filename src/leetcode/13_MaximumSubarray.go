package main

import "fmt"

func maxSubArray(nums []int) int {
	data := make([][2]int, len(nums))
	for i:= 0; i < len(nums); i++ {
		data[i][0] = nums[i]
	}
	return compute(data, len(nums) - 1)
}

func compute(data [][2]int, n int) int{
	if n == 0 {
		data[0][1] = data[0][0]
		return data[0][0]
	}
	maxLastNum := compute(data, n - 1)
	diejia1 := data[n-1][1] + data[n][0]
	max := max(maxLastNum, diejia1, data[n][0],data,n)
	return max
}
func max(maxLastNum,diejia,currentVal int, data [][2]int, num int) int {
	if maxLastNum >= diejia && maxLastNum >= currentVal {//maxLastNum
		if diejia < currentVal {
			data[num][1] = currentVal
		} else {
			data[num][1] = diejia
		}
		return maxLastNum
	} else if diejia >= maxLastNum && diejia >= currentVal {//diejia
		data[num][1] = diejia
		return diejia
	} else if currentVal >= diejia && currentVal >= maxLastNum {//currentVal
		data[num][1] = currentVal
		return currentVal
	} else {
		fmt.Println("出问题了######")
		fmt.Printf("%d,%d,%d", maxLastNum,diejia,currentVal)
		fmt.Println("完######")
		return 0
	}
}


func main() {
	fmt.Println(maxSubArray([]int{8,-19,5,-4,20}))
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
	fmt.Println(maxSubArray([]int{1,-1,1}))
}
