//package main
////
////var a = "G"  // 全局变量
////
////// GOG
////func main() {
////	n()
////	m()
////	n()
////}
////
////func n() { print(a) }
////
////func m() {
////	a := "O"	// 局部变量
////	print(a)
////}
//
//package main
//
//var a = "G"
//
//// GOO
//func main() {
//	n()
//	m()
//	n()
//}
//
//func n() {
//	print(a)
//}
//
//func m() {
//	a = "O"
//	print(a)
//}
package main

var a string

// GOG
func main() {
	a = "G"
	print(a)
	f1()
}

func f1() {
	a := "O"
	print(a)
	f2()
}

func f2() {
	print(a)
}
