package main

/*
给你一个数组 nums ，数组中有 2n 个元素，按 [x1,x2,...,xn,y1,y2,...,yn] 的格式排列。

请你将数组按 [x1,y1,x2,y2,...,xn,yn] 格式重新排列，返回重排后的数组

1 <= n <= 500
nums.length == 2n
1 <= nums[i] <= 10^3
*/

func shuffle(nums []int, n int) []int {
	var res []int
	left := n
	for i := 0; i < n; i++ {
		res = append(res, nums[i])
		res = append(res, nums[left])
		left += 1
	}
	return res
}
