package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 测试字符串删除字符
	a := "12321312"
	fmt.Println(strings.Replace(a, "5", "", 1))

	// 测试整数转二进制
	fmt.Println(strconv.FormatInt(9, 2))
}
