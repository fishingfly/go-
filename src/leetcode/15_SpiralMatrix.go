package main

import (
	"fmt"
	"strconv"
)


func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 {
		return nil
	}
	result := make([]int, len(matrix) * len(matrix[0]))
	var posMap = make(map[string]string)
	i := 0
	j := -1
	index := -1
	for {
		//顺时针开始旋转吧
		i,j,index = goRight(matrix, i, j, result, index,posMap)
		i,j,index = goDown(matrix, i, j, result, index,posMap)
		i,j,index = goLeft(matrix, i, j, result, index,posMap)
		i,j,index = goUp(matrix, i, j, result, index,posMap)
		//什么时候跳出呢
		if len(posMap) == len(matrix) * len(matrix[0]) {
			break
		}
	}
	return result
}

//right
func goRight(matrix [][]int, i,j int,result []int, index int,posMap map[string]string) (int, int,int){
	indexCol := j
	indexCol++
	for ; indexCol <  len(matrix[i]); indexCol++ {//行不动，列动
		_, ok := posMap[strconv.Itoa(i)+strconv.Itoa(indexCol)]
		if ok {
			return i,indexCol - 1,index
		} else {
			index++
			result[index] = matrix[i][indexCol]
			posMap[strconv.Itoa(i)+strconv.Itoa(indexCol)] = "1"
		}
	}
	return i,indexCol - 1,index
}

//down
func goDown(matrix [][]int, i,j int,result []int, index int,posMap map[string]string) (int, int,int){
	indexRow := i
	indexRow++
	for ; indexRow <  len(matrix); indexRow++ {//行动，列不动
		_, ok := posMap[strconv.Itoa(indexRow)+strconv.Itoa(j)]
		if ok {
			return indexRow - 1,j,index
		} else {
			index++
			result[index] = matrix[indexRow][j]
			posMap[strconv.Itoa(indexRow)+strconv.Itoa(j)] = "1"
		}
	}
	return indexRow - 1,j,index
}

//left
func goLeft(matrix [][]int, i,j int,result []int, index int,posMap map[string]string) (int,int,int) {
	indexCol := j
	indexCol--
	for ; indexCol >= 0; indexCol-- {//行不动，列动
		_, ok := posMap[strconv.Itoa(i)+strconv.Itoa(indexCol)]
		if ok {
			return i,indexCol + 1,index
		} else {
			index++
			result[index] = matrix[i][indexCol]
			posMap[strconv.Itoa(i)+strconv.Itoa(indexCol)] = "1"
		}
	}
	return i,indexCol+1,index
}

//up
func goUp(matrix [][]int, i,j int,result []int, index int,posMap map[string]string) (int,int,int){
	indexRow := i
	indexRow--
	for ; indexRow >= 0; indexRow-- {//行动，列不动
		_, ok := posMap[strconv.Itoa(indexRow)+strconv.Itoa(j)]
		if ok {
			return indexRow + 1,j,index
		} else {
			index++
			result[index] = matrix[indexRow][j]
			posMap[strconv.Itoa(indexRow)+strconv.Itoa(j)] = "1"
		}
	}
	return indexRow+1,j,index
}


func main() {
	fmt.Println(spiralOrder([][]int{
		 {1, 2, 3 },
	{4, 5, 6 },
	{7, 8, 9 },
	}))
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3, 4 },
		{5, 6, 7, 8 },
		{9,10,11,12 },
	}))
	fmt.Println(spiralOrder([][]int{
	}))
}
