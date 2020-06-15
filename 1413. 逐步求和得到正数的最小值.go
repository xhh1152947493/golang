package main

import "fmt"

/*
给你一个整数数组 nums 。你可以选定任意的 正数 startValue 作为初始值。

你需要从左到右遍历 nums 数组，并将 startValue 依次累加上 nums 数组中的值。

请你在确保累加和始终大于等于 1 的前提下，选出一个最小的 正数 作为 startValue 。

1 <= nums.length <= 100
-100 <= nums[i] <= 100
*/

func minStartValue(nums []int) int {
	l := len(nums)
	res := 1
	for true {
		flag := false
		tmp := res
		for i, v := range nums {
			tmp += v
			if tmp < 1 {
				break
			}
			if i == l-1 {
				flag = true
			}
		}
		if flag {
			return res
		}
		res += 1
	}
	return -1
}

func main() {
	var test = []int{-3, 2, -3, 4, 2}
	fmt.Println(minStartValue(test))
}
