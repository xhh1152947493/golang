package main

/*
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。
1 <= nums.length <= 50000
1 <= nums[i] <= 10000
*/

func exchange(nums []int) []int {
	var arr1 []int
	var arr2 []int
	for _, num := range nums {
		if num%2 == 0 {
			arr2 = append(arr2, num)
		} else {
			arr1 = append(arr1, num)
		}
	}
	return append(arr1, arr2...) // 列表追加，...表示列表所有元素
}
