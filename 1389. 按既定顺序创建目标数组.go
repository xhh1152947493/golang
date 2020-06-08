package main

/*
给你两个整数数组 nums 和 index。你需要按照以下规则创建目标数组：

目标数组 target 最初为空。
按从左到右的顺序依次读取 nums[i] 和 index[i]，在 target 数组中的下标 index[i] 处插入值 nums[i] 。
重复上一步，直到在 nums 和 index 中都没有要读取的元素。
请你返回目标数组。

题目保证数字插入位置总是存在
1 <= nums.length, index.length <= 100
nums.length == index.length
0 <= nums[i] <= 100
0 <= index[i] <= i

*/

func createTargetArray(nums []int, index []int) []int {
	var res []int
	for i, v := range nums {
		if i == index[i] {
			res = append(res, v)
		} else {
			res = append(res, v) // 先将元素加入，然后移动元素
			for j := i; j > index[i]; j-- {
				res[j] = res[j-1]
			}
			res[index[i]] = v
		}
	}
	return res
}

func main() {
	var a = []int{0, 1, 2, 3, 4}
	b := []int{0, 1, 2, 2, 1}
	createTargetArray(a, b)
}
