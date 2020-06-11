package main

/*
给你一个整数 n，请你返回 任意 一个由 n 个 各不相同 的整数组成的数组，并且这 n 个数相加和为 0 。

1 <= n <= 1000
*/

func sumZero(n int) []int {
	var res []int
	flag := n%2 == 1
	l := n / 2
	for i := n; i > n-l; i-- {
		res = append(res, i)
		res = append(res, -i)
	}
	if flag {
		res = append(res, 0)
	}
	return res
}
