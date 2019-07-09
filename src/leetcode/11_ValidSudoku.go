package main

func isValidSudoku(board [][]byte) bool {
//收集一点横竖的数字排除掉，周围九宫格的数字排除，还剩多少
//最暴力的方式是遍历
//回溯条件是这个格子无任何数字可放
//我觉得是用递归做的


//速度快的方法是，先去找可填数字最少的各自，从这些格子开始向外扩展，其实就是剪枝。这个会快点
}

//返回当前格子的可填数字
func getNumListForCurrentCell(x,y int) []int {
	list := make([]int, 9)
	for i := 0; i< 9; i++ {
		list[i] = 1
	}
	//横
	for ()
	//竖
	//周围九宫格
}

func main() {

}


