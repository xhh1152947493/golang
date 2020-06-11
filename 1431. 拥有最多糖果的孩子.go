package main

/*
给你一个数组 candies 和一个整数 extraCandies ，其中 candies[i] 代表第 i 个孩子拥有的糖果数目。

对每一个孩子，检查是否存在一种方案，将额外的 extraCandies 个糖果分配给孩子们之后，此孩子有 最多 的糖果。注意，允许有多个孩子同时拥有 最多 的糖果数目。

2 <= candies.length <= 100
1 <= candies[i] <= 100
1 <= extraCandies <= 50
*/

func kidsWithCandies(candies []int, extraCandies int) []bool {
	// 找到最大值
	maxVal := candies[0]
	for i := 1; i < len(candies); i++ {
		if maxVal < candies[i] {
			maxVal = candies[i]
		}
	}
	var res []bool
	for _, v := range candies {
		if extraCandies+v >= maxVal {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res
}
