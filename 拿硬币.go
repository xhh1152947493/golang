package main

func minCount(coins []int) int {
	res := 0
	for _, v := range coins {
		if v%2 == 1 {
			res += v/2 + 1
		} else {
			res += v / 2
		}
	}
	return res
}
