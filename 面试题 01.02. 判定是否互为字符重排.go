package main

import "strings"

/*
给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

0 <= len(s1) <= 100
0 <= len(s2) <= 100

题解思路：
1、遍历s1，对于s1中的每一个元素x，判断是否在s2中，如果在则删除s2中的x
2、遍历完成后，如果s2为空，则为true
*/

func CheckPermutation(s1 string, s2 string) bool {
	for _, v := range s1 {
		s2 = strings.Replace(s2, string(v), "", 1) // 字符串删除用Replace
	}
	if len(s2) > 0 {
		return false
	}
	return true
}
