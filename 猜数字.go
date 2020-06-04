package main

import "fmt"

func game(guess []int, answer []int) int {
	var res = 0
	for i := 0; i < len(guess); i++ {
		if guess[i] == answer[i] {
			res += 1
		}
	}
	return res
}

func main() {
	guess := []int{1, 2, 3}
	answer := []int{1, 3, 2}
	fmt.Println(game(guess, answer))
}
