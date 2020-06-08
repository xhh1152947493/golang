package main

import "fmt"

func replaceElements(arr []int) []int {
	arr[len(arr)-1] = -1
	for i := range arr[:len(arr)-1] {
		_, rightMax := MaxInt(arr[i:]) // 要调用MaxInt，必须同时编译
		if arr[i] < rightMax {
			arr[i] = rightMax
		}
	}
	return arr
}

func main() {
	var balance = []int{1000, 2, -1, 3, 17, 50, 0, -1, -1}
	args := replaceElements(balance)
	fmt.Println(args)
}
