package main

/*
给定一个包含 0 和 1 的二维网格地图，其中 1 表示陆地 0 表示水域。

网格中的格子水平和垂直方向相连（对角线方向不相连）。整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。

岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿的周长。

题解思路：
1、遍历二维数组，对于每个陆地1，遍历其4个方向。如果为水域0则周长+1
*/

func islandPerimeter(grid [][]int) int {
	width := len(grid)     // 宽
	height := len(grid[0]) // 长
	res := 0
	for i, row := range grid {
		for j, col := range row {
			if col == 1 {
				res += findAreaWater(i, j, width, height, grid)
			}
		}
	}
	return res
}

func findAreaWater(i int, j int, width int, height int, grid [][]int) int {
	// 找到周围的水域
	tmp := 0
	if checkInArea(i, j-1, width, height) {
		if grid[i][j-1] == 0 {
			tmp += 1
		}
	} else {
		tmp += 1
	}

	if checkInArea(i-1, j, width, height) {
		if grid[i-1][j] == 0 {
			tmp += 1
		}
	} else {
		tmp += 1
	}

	if checkInArea(i, j+1, width, height) {
		if grid[i][j+1] == 0 {
			tmp += 1
		}
	} else {
		tmp += 1
	}

	if checkInArea(i+1, j, width, height) {
		if grid[i+1][j] == 0 {
			tmp += 1
		}
	} else {
		tmp += 1
	}
	return tmp
}

func checkInArea(i int, j int, width int, height int) bool {
	// 检查该点是否在范围内
	return i >= 0 && i < width && j >= 0 && j < height
}
