package main

import "strings"

/*
给你一个整数 n，请你返回一个含 n 个字符的字符串，其中每种字符在该字符串中都恰好出现 奇数次 。

返回的字符串必须只含小写英文字母。如果存在多个满足题目要求的字符串，则返回其中任意一个即可

1 <= n <= 500

*/

func generateTheString(n int) string {
	//a := "x"
	//b := "y"
	//if n%2 == 1 {
	//	res := ""
	//	for i := 0; i < n; i++ {
	//		res += a
	//	}
	//	return res
	//} else {
	//	res := ""
	//	for i := 0; i < n-1; i++ {
	//		res += a
	//	}
	//	res += b
	//	return res
	//}

	if n&1 == 1 {
		return strings.Repeat("a", n)
	}

	return strings.Repeat("a", n-1) + "b"

}
