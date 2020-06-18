package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
给定一个字符串S，通过将字符串S中的每个字母转变大小写，我们可以获得一个新的字符串。返回所有可能得到的字符串集合。

S 的长度不超过12。
S 仅由数字和字母组成。
*/

func letterCasePermutation(S string) []string {
	aAscii, _ := strconv.Atoi("a")
	zAscii, _ := strconv.Atoi("z")
	AAscii, _ := strconv.Atoi("A")
	ZAscii, _ := strconv.Atoi("Z")
	for _, v := range S {
		if int(v) >= aAscii && int(v) <= zAscii {
			S = strings.Replace(S, string(v), string(v+96), 1)
		} else if int(v) >= AAscii && int(v) <= ZAscii {
			S = strings.Replace(S, string(v), string(v-96), 1)
		}

	}
	var res []string
	return res
}

func main() {
	var test = "a1b2"
	fmt.Println(letterCasePermutation(test))
}
