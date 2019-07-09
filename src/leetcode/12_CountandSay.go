package main

import (
	"fmt"
	"strconv"
)

var numToStrMap = make(map[int]string)


func countAndSay(n int) string {
	if n < 1 || n > 30 {
		return "error"
	}
	// 先搜索map中没有该数字
	val ,ok := numToStrMap[n]
	if ok {
		return val
	}
	if n == 1 {
		return "1"
	}
	//进入搜索小一个数字的排列
	smallNumStr := countAndSay(n - 1)
	numToStrMap[n] = numToStr(smallNumStr)
	return numToStrMap[n]
}

func numToStr(numStr string) string{
	pinjie := ""
	count := 0
	if len(numStr) == 1 {
		pinjie += "1"
		pinjie += string(numStr[0])
		return pinjie
	}
	for i := 0; i < len(numStr); i++ {
		if count == 0 {
			count++
			continue
		}
		if numStr[i] == numStr[i - 1] {
			count++
		} else {
			pinjie += strconv.Itoa(count)
			pinjie += string(numStr[i- 1])
			count = 1
		}
		if i == len(numStr) - 1 {
			pinjie += strconv.Itoa(count)
			pinjie += string(numStr[i])
		}
	}
	return pinjie
}

func main() {
	//numToStr("11")// 21
	//numToStr("21")// 1211
	//numToStr("1211")// 111221
	//numToStr("111221")// 312211
	fmt.Println(countAndSay(5))


}