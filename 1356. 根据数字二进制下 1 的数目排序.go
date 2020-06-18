package main

import (
	"fmt"
	"strconv"
)

/*
给你一个整数数组 arr 。请你将数组中的元素按照其二进制表示中数字 1 的数目升序排序。

如果存在多个数字二进制中 1 的数目相同，则必须将它们按照数值大小升序排列。

请你返回排序后的数组

1 <= arr.length <= 500
0 <= arr[i] <= 10^4

题解思路：
1、遍历arr，对于每一个元素x转换成二进制
2、统计1的个数，并映射到map中
3、按个数排序
*/

func sortByBits(arr []int) []int {
	_map := make(map[int]int)
	//var res [] int
	for _, v := range arr {
		_map[v] = 0
		bitStr := strconv.FormatInt(int64(v), 2) // 转成二进制字符串
		for _, b := range bitStr {
			if b == 49 { // ASCII 为49，即整形1
				_map[v] += 1 // 统计每个元素对应二进制1的个数
			}
		}
	}

	// 冒泡排序
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if _map[arr[i]] > _map[arr[j]] || (_map[arr[i]] == _map[arr[j]] && arr[i] > arr[j]) {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	return arr

	// 插入排序
	//var res = []int{arr[0]}
	//for k, v := range _map{
	//
	//}

	//for k, v := range _map {
	//	if len(res) > 0 {
	//		for i, num := range res {
	//			if _map[num] > v {
	//				if i == len(res)-1 {
	//					res = append(res, k)
	//				} else {
	//					continue
	//				}
	//			} else if _map[num] < v {
	//				res[i+1], res[i] = res[i], _map[num] // 插入到前面
	//			} else {
	//				if k >= num {
	//					res = append(res, k)
	//				} else {
	//					res[i+1], res[i] = res[i], _map[num] // 插入到前面
	//				}
	//			}
	//		}
	//	} else {
	//		res = append(res, k)
	//	}
	//}
	//return res
}

func main() {
	var test = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(sortByBits(test))
}
