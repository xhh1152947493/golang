package main

import "fmt"

/*
给你一个整数 n 。请你先求出从 1 到 n 的每个整数 10 进制表示下的数位和（每一位上的数字相加），然后把数位和相等的数字放到同一个组中。

请你统计每个组中的数字数目，并返回数字数目并列最多的组有多少个。

1 <= n <= 10^4

题解思路：
1、求每个整数的数位和，并映射到map中
*/

func countLargestGroup(n int) int {
	_map := make(map[int]int)
	countMap := make(map[int]int)
	_max := 0
	res := 0
	for i := 1; i <= n; i++ {
		tmpSum := 0
		tmpWei := i
		flag := true
		for flag {
			if tmpWei >= 10 {
				tmpSum += tmpWei % 10
				tmpWei /= 10
			} else {
				tmpSum += tmpWei
				flag = false
			}
		}

		//if tmpWei >= 10 {
		//	for tmpWei >= 10 {
		//		tmpSum += tmpWei % 10
		//		tmpWei /= 10
		//	}
		//} else {
		//	tmpSum = i
		//}
		//for j := 1; j <= i; j++ { // 求数位和
		//	tmpSum += j
		//}
		_map[i] = tmpSum
	}
	for _, v := range _map { // 统计相同数位和整数的个数
		countMap[v] += 1
	}
	for _, v := range countMap { // 找到最大值
		if v > _max {
			_max = v
		}
	}
	for _, v := range countMap { // 找到最大值的个数
		if v == _max {
			res += 1
		}
	}
	return res
}

func main() {
	var test = 15
	fmt.Println(countLargestGroup(test))
}
