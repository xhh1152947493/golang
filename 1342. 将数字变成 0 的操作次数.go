package main

/*
给你一个非负整数 num ，请你返回将它变成 0 所需要的步数。 如果当前数字是偶数，你需要把它除以 2 ；否则，减去 1 。

0 <= num <= 10^6

*/

func numberOfSteps(num int) int {
	res := 0
	for num != 0 {
		res += 1
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
	}
	return res
}
