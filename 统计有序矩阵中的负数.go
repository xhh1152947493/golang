package main

import "fmt"

/*
给你一个 m * n 的矩阵 grid，矩阵中的元素无论是按行还是按列，都以非递增顺序排列。

请你统计并返回 grid 中 负数 的数目。
m == grid.length
n == grid[i].length
1 <= m, n <= 100
-100 <= grid[i][j] <= 100
*/
func countNegatives(grid [][]int) int {
	res := 0
	for _, value1 := range grid {
		for i, value2 := range value1 {
			if value2 < 0 { // 有序数组直接加
				res += len(value1) - i
				break
			}
		}
	}
	return res
}

func main() {
	var a = [][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{1, 1, -1, -2},
		{-1, -1, -2, -3},
	}
	fmt.Println(countNegatives(a))
}
