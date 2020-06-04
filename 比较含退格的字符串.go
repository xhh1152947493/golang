package main

import "fmt"

func backspaceCompare(S string, T string) bool {
	stack := make([]rune, 0, len(S))
	for _, c := range S {
		if c == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, c)
		}
	}
	s := string(stack)
	stack = stack[:0]
	for _, c := range T {
		if c == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, c)
		}
	}
	t := string(stack)
	return s == t
}

func main() {
	fmt.Println(backspaceCompare("1#23", "4#56"))
}
