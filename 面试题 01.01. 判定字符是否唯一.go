package main

import "reflect"

/*
实现一个算法，确定一个字符串 s 的所有字符是否全都不同。

0 <= len(s) <= 100
如果你不使用额外的数据结构，会很加分。
*/

func isUnique(astr string) bool {
	var exist []string
	for _, v1 := range astr {
		for _, v2 := range exist {
			if reflect.DeepEqual(v1, v2) {
				return false
			}
		}
		exist = append(exist, v1)
	}
	return true
}
