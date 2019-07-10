package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	begin := -1
	end := -1
	result := make([][]int,0, len(intervals)+1)
	if len(intervals) == 0 {
		result = append(result, []int{newInterval[0], newInterval[1]})
		return result
	}
	for i := 0; i < len(newInterval); i++ {
		for j := 0; j < len(intervals); j++ {
			if newInterval[i] >= intervals[j][0] && newInterval[i] <= intervals[j][1] {//在区间中
				if i == 0 {
					begin = intervals[j][0]
					break
				} else {
					end = intervals[j][1]
				}
			} else if newInterval[i] < intervals[j][0] {// 在区间左侧
				if i == 0 {
					begin = newInterval[i]
					break
				} else {
					if end == -1 {
						end = newInterval[i]
					}
					if begin != -1 && end != -1 {
						result = append(result, []int{begin, end})
						begin,end = -1,-1
					}
					result = append(result, intervals[j])
				}
			} else if newInterval[i] > intervals[j][1] {// 在区间右侧
				if i == 0 {
					result = append(result, intervals[j])
					begin = newInterval[i]
				} else {
					end = newInterval[i]
				}
			}
		}
	}
	if begin != -1 && end != -1 {
		result = append(result, []int{begin, end})
	}
	return result
}
func main() {
	fmt.Println(insert([][]int{
		{1,2},{3,5},{6,7},{8,10},{12,16},
	}, []int{4,8}))
	fmt.Println(insert([][]int{
		{0,6},{8,8},{9,15},{16,21},
	}, []int{1,2}))
	fmt.Println(insert([][]int{
		{0,6},{8,8},{9,15},{16,21},
	}, []int{8,22}))
	fmt.Println(insert([][]int{
	}, []int{8,22}))
	fmt.Println(insert([][]int{{1,5},
	}, []int{0,0}))
	fmt.Println(insert([][]int{{1,2},{3,5},{6,7},{8,10},{12,16},
	}, []int{4,8}))

}
