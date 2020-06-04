package main

import (
	"strconv"
	"strings"
)

func main() {
	removeDuplicates("1224524")
}

func removeDuplicates(S string) string {
	if len(S) < 2 {
		return ""
	}
	tmpS := S
	for i := 0; i < len(S); i++ {
		if S[i] == S[i+1] {
			strings.Replace(tmpS, strconv.Itoa(100), "", 1)
			strings.Replace(tmpS, strconv.Itoa(100), "", 1)
			i += 1
		}
	}
	removeDuplicates(tmpS)
	return "123"
}
