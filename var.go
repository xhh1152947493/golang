package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}

// 声明变量的一般形式是使用 var 关键字：var identifier type。
/*这种因式分解关键字的写法一般用于声明全局变量。
var (
    a int
    b bool
    str string
)
*/

// 变量的命名规则遵循骆驼命名法，即首个单词小写，每个新单词的首字母大写，例如：numShips 和 startDate。
// 原因，首字母大写在包外可见，小写不可见。可见性规则

/* nb，直接在编译时就完成推断变量类型的过程，相当于将强类型转换成弱类型？
声明与赋值（初始化）语句也可以组合起来。

示例：

var identifier [type] = value
var a int = 15
var i = 5
var b bool = false
var str string = "Go says hello to the world!"
但是 Go 编译器的智商已经高到可以根据变量的值来自动推断其类型，这有点像 Ruby 和 Python 这类动态语言，只不过它们是在运行时进行推断，而 Go 是在编译时就已经完成推断过程。因此，你还可以使用下面的这些形式来声明及初始化变量：

var a = 15
var b = false
var str = "Go says hello to the world!"
或：

var (
    a = 15
    b = false
    str = "Go says hello to the world!"
    numShips = 50
    city string
)
*/

/*函数体外一般为全局变量，在函数体外声明局部变量用 ：=  局部变量在循环和判断时可以用
这种写法主要用于声明包级别的全局变量，当你在函数体内声明局部变量时，应使用简短声明语法 :=，例如：

a := 1
*/

/*
a, b, c = 5, 7, "abc"
上面这行假设了变量 a，b 和 c 都已经被声明，否则的话应该这样使用：

a, b, c := 5, 7, "abc"
右边的这些值以相同的顺序赋值给左边的变量，所以 a 的值是 5， b 的值是 7，c 的值是 "abc"。

这被称为 并行 或 同时 赋值。

如果你想要交换两个变量的值，则可以简单地使用 a, b = b, a。

(在 Go 语言中，这样省去了使用交换函数的必要)

空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。
*/
