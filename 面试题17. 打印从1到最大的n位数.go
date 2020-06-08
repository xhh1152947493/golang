package main

import (
	"fmt"
	"math"
)

func printNumbers(n int) []int {
	l := math.Pow(10, float64(n)) // 最大长度
	var res []int                 // 变长切片
	for i := 1; i < int(l); i++ {
		res = append(res, i)
	}
	return res
}

func main() {
	fmt.Println(printNumbers(1))
}
