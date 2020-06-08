package main

import "sort"

/*
给你一个整数数组 nums，请你选择数组的两个不同下标 i 和 j，使 (nums[i]-1)*(nums[j]-1) 取得最大值。

请你计算并返回该式的最大值

2 <= nums.length <= 500
1 <= nums[i] <= 10^3
*/

func maxProduct(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	return (nums[l-1] - 1) * (nums[l-2] - 1)
}
