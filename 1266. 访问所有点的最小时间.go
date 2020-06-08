package main

import (
	"fmt"
	"math"
)

func minTimeToVisitAllPoints(points [][]int) int {
	res := 0
	L := len(points)
	if L <= 1 {
		return 0
	}
	for index := range points[:L-1] { //切片
		start_ := points[index]
		end_ := points[index+1]
		y_distens := end_[0] - start_[0]
		x_distens := end_[1] - start_[1]
		if x_distens == y_distens {
			if x_distens < 0 {
				res += -x_distens
			} else {
				res += x_distens
			}
		} else {
			res += int(math.Sqrt(math.Pow(float64(end_[0]-start_[0]), 2) + math.Pow(float64(end_[1]-start_[1]), 2)))
		}
	}
	return res
}

func main() {
	var arr = [][]int{
		{3, 2},
		{-2, 2},
	}
	fmt.Println(minTimeToVisitAllPoints(arr))
}
