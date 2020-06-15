package main

import (
	"fmt"
	"sort"
)

/*
给你一个整数数组 A，对于每个整数 A[i]，我们可以选择处于区间 [-K, K] 中的任意数 x ，将 x 与 A[i] 相加，结果存入 A[i] 。

在此过程之后，我们得到一些数组 B。

返回 B 的最大值和 B 的最小值之间可能存在的最小差值

1 <= A.length <= 10000
0 <= A[i] <= 10000
0 <= K <= 10000

*/

func smallestRangeI(A []int, K int) int {
	/*
		求平均值，然后每个元素在K这个区间找一个值靠近这个值
	*/
	sum_ := 0
	l := len(A)
	for _, v := range A {
		sum_ += v
	}
	ave := sum_ / l
	for i, v := range A {
		if v < ave {
			for j := K; j > 0; j-- {
				if j+v <= ave {
					A[i] = j + v
					break
				}
			}
		} else if v > ave {
			for j := -K; j < 0; j++ {
				if j+v >= ave {
					A[i] = j + v
					break
				}
			}
		} else {
			continue
		}
	}
	sort.Ints(A)
	return A[len(A)-1] - A[0]
}

func main() {
	var test = []int{3, 1, 10}
	K := 4
	fmt.Println(smallestRangeI(test, K))
}
