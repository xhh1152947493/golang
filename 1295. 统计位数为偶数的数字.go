package main

/*
给你一个整数数组 nums，请你返回其中位数为 偶数 的数字的个数。
1 <= nums.length <= 500
1 <= nums[i] <= 10^5
*/
func findNumbers(nums []int) int {
	res := 0
	for _, v := range nums {
		flag := true
		tmp_ := 0
		for flag {
			v /= 10
			tmp_ += 1
			if v == 0 {
				flag = false
			}
		}
		if tmp_ != 0 && tmp_%2 == 0 {
			res += 1
		}
	}
	return res
}

/* python代码一行,但是时间上的效率低了很多
return sum([1 if len(str(num)) % 2 == 0 else 0 for num in nums ])
*/
