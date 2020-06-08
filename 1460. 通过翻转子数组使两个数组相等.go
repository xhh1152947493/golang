package main

import (
	"reflect"
	"sort"
)

func canBeEqual(target []int, arr []int) bool {
	sort.Ints(target) // 整形数组排序
	sort.Ints(arr)
	return reflect.DeepEqual(target, arr) // DeepEqual，判断是否相等。不能用python中的== 直接判断
}
