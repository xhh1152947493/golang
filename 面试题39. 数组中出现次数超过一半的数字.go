package main

import "sort"

/*
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

1 <= 数组长度 <= 50000

*/

func majorityElement(nums []int) int {
	sort.Ints(nums) // sort.Ints 无返回值
	return nums[len(nums)/2]
}
