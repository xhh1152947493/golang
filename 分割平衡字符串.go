package main

import "fmt"

func balancedStringSplit(s string) int {
	if len(s) < 2 {
		return 0
	}

	res := 0
	count := 1
	if s[0] == 'R' {
		count = -1
	}

	for i := 1; i < len(s); i++ {
		if 'R' == s[i] {
			count--
		} else if 'L' == s[i] {
			count++
		}
		if count == 0 {
			res++
		}
	}

	return res
}

func main() {
	fmt.Println(balancedStringSplit("LRRLRLLLR"))
}
