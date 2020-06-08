package main

/*
给你一个整数 n，请你帮忙计算并返回该整数「各位数字之积」与「各位数字之和」的差。

1 <= n <= 10^5
*/

func subtractProductAndSum(n int) int {
	sum_ := 0
	mut_ := 1
	for n != 0 {
		lastNum := n % 10
		n /= 10
		sum_ += lastNum
		mut_ *= lastNum
	}
	return mut_ - sum_
}

func main() {
	subtractProductAndSum(14)
}
