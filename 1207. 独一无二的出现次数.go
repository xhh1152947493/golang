package main

import (
	"fmt"
)

/*
给你一个整数数组 arr，请你帮忙统计数组中每个数的出现次数。

如果每个数的出现次数都是独一无二的，就返回 true；否则返回 false。

1 <= arr.length <= 1000
-1000 <= arr[i] <= 1000

*/
func uniqueOccurrences(arr []int) bool {
	//len_ := len(arr)
	map_ := make(map[int]int)
	//var tar [] int
	//for i := 0; i < len_; i++ {
	//	if _, ok := map_[arr[i]]; ok {
	//		map_[arr[i]] += 1
	//	} else {
	//		map_[arr[i]] = 1
	//	}
	//}

	// 以下代码等价上述代码
	for _, v := range arr {
		map_[v]++
	}
	vFlag := make(map[int]bool)
	for _, v := range map_ {
		if vFlag[v] { // 如果value已经访问过，则设置访问标志为true，再次访问到时直接返回false
			return false
		} else {
			vFlag[v] = true
		}
	}
	return true

	//for _, v := range map_ {
	//	tar = append(tar, v)
	//}
	//for i := 0; i < len(tar)-1; i++ {
	//	for j := i + 1; j < len(tar); j++ {
	//		if tar[j] == tar[i] {
	//			return false
	//		}
	//	}
	//}
	//return true
}

func main() {
	var test = []int{1, 2, 2, 1, 1, 3}
	fmt.Println(uniqueOccurrences(test))
}
