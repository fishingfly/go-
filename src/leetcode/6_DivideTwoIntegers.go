package main

import (
	"fmt"
	"math"
)

//我是用减法做的，用左移会快点,不能用乘除法就要想到用位运算
func divide(dividend int, divisor int) int {
	divisorCount := 0
	absDivisor := math.Abs(float64(divisor))
	absDividend := math.Abs(float64(dividend))
	if absDividend < absDivisor {
		return 0
	}
	if absDivisor != 1 {
		for {
			if absDividend == 0 {
				break
			} else {
				if (absDividend - absDivisor) < 0 {
					absDividend = 0
				} else {
					absDividend -= absDivisor
					divisorCount++
				}
			}
		}
	} else {
		divisorCount = int(absDividend)
	}
	if !isPositive(dividend, divisor) {
		divisorCount = 0 - divisorCount
	}
	return afterProcess(divisorCount)
}

func afterProcess(num int) int{
	if num > 2147483647  || num < -2147483648{
		num = 2147483647
	}
	return num
}

func isPositive(dividend int, divisor int) bool {
	if dividend > 0 {
		if divisor > 0 {
			return true
		} else {
			return false
		}
	} else if dividend < 0 {
		if divisor > 0 {
			return false
		} else {
			return true
		}
	} else {
		return true
	}
}

func main() {
	fmt.Println(5<<2)
	//fmt.Println(divide(-999511578,-2147483648))
}