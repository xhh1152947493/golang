package main

import (
	"fmt"
	"strings"
)

func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20") // strings.ReplaceAll替换所有符合条件的字符
}

func main() {
	str_ := "We are happy."
	str_ = replaceSpace(str_)
	fmt.Println(str_)
}
