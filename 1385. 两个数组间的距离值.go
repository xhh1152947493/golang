package main

import "math"

/*
给你两个整数数组 arr1 ， arr2 和一个整数 d ，请你返回两个数组之间的 距离值 。

「距离值」 定义为符合此描述的元素数目：对于元素 arr1[i] ，不存在任何元素 arr2[j] 满足 |arr1[i]-arr2[j]| <= d 。

1 <= arr1.length, arr2.length <= 500
-10^3 <= arr1[i], arr2[j] <= 10^3
0 <= d <= 100

*/

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res := 0
	for _, v1 := range arr1 {
		flag := true
		for _, v2 := range arr2 {
			if math.Abs(float64(v1-v2)) <= float64(d) {
				flag = false
				break
			}
		}
		if flag {
			res += 1
		}
	}
	return res
}
