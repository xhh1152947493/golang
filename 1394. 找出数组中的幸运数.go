package main

import (
	"fmt"
	"sort"
)

/*
在整数数组中，如果一个整数的出现频次和它的数值大小相等，我们就称这个整数为「幸运数」。

给你一个整数数组 arr，请你从中找出并返回一个幸运数。

如果数组中存在多个幸运数，只需返回 最大 的那个。
如果数组中不含幸运数，则返回 -1

1 <= arr.length <= 500
1 <= arr[i] <= 500
*/

func findLucky(arr []int) int {
	sort.Ints(arr)
	len_ := len(arr)
	map_ := make(map[int]int)
	//map_[1] = 2
	//count := 1
	//res := -1
	for i := 0; i < len_; i++ {
		//a, b := map_[i] // a为map_[i]的值，b为是否存在
		//fmt.Println(a, b)
		if _, ok := map_[arr[i]]; ok {
			map_[arr[i]] += 1
		} else {
			map_[arr[i]] = 1
		}
	}
	for i := len_ - 1; i >= 0; i-- {
		if map_[arr[i]] == arr[i] {
			return arr[i]
		}
	}
	return -1
}

func main() {
	var test = []int{19, 12, 11, 14, 18, 8, 6, 6, 13, 9, 8, 3, 10, 10, 1, 10, 5, 12, 13, 13, 9}
	fmt.Println(findLucky(test))
}
