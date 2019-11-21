package main

import "fmt"

/*
常量使用关键字 const 定义，用于存储不会改变的数据。

存储在常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。

常量的定义格式：const identifier [type] = value，例如：

const Pi = 3.14159
在 Go 语言中，你可以省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型。

显式类型定义： const b string = "abc"
隐式类型定义： const b = "abc"
*/

func main() {
	const Pi = 3.14159
	//const b = "abc" // => const b string = "abc"
	//const beef, two, c = "eat", 2, "veg"
	// => const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6
	const (
		Monday, Tuesday, Wednesday = 1, 2, 3
		Thursday, Friday, Saturday = 4, 5, 6
	)
	/* 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：
	const (
	    a = iota
	    b
	    c
	)
	*/
	const (
		a = iota
		b = iota
		c = iota
	)
	//const d = 100 + iota
	// iota 也可以用在表达式中，如：iota + 50。在每遇到一个新的常量块或单个常量声明时， iota 都会重置为 0（ 简单地讲，每遇到一次 const 关键字，iota 就重置为 0 ）。
	fmt.Println(a, b, c, iota+100)
}
