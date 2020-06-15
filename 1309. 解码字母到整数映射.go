package main

import (
	"fmt"
	"strconv"
)

/*
给你一个字符串 s，它由数字（'0' - '9'）和 '#' 组成。我们希望按下述规则将 s 映射为一些小写英文字符：

字符（'a' - 'i'）分别用（'1' - '9'）表示。
字符（'j' - 'z'）分别用（'10#' - '26#'）表示。
返回映射之后形成的新字符串。

题目数据保证映射始终唯一。

1 <= s.length <= 1000
s[i] 只包含数字（'0'-'9'）和 '#' 字符。
s 是映射始终存在的有效字符串。
*/

func freqAlphabets(s string) string {
	var ans []rune // int32类型
	n := len(s)
	for i := 0; i < n; i++ {
		if i+2 < n && s[i+2] == '#' {
			ans = append(ans, get(s[i:i+2]))
			i += 2
		} else {
			ans = append(ans, get(s[i:i+1]))
		}
	}
	return string(ans) // ascii数值转字符串
}

func get(s string) rune {
	r, _ := strconv.Atoi(s) // 字符串转数值
	return rune(r + 96)     // 转ascii码
}

func main() {
	res := freqAlphabets("10#11#12")
	fmt.Println(res)
}
