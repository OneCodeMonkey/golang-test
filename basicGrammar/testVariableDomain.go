package main

import (
	"fmt"
)

var a int = 13


func func1() {
	var a int = 26

	fmt.Printf("%d\n", a)
}


func func2(a int) {
	a = 14

	fmt.Printf("%d\n", a)
}


func main() {
	fmt.Printf("%d\n", a)

	func1();

	fmt.Printf("%d\n", a)	// 13.  验证：局部定义的变量，不会覆盖全局定义的同名变量

	func2(a)

	fmt.Printf("%d\n", a)	// 13. 验证：函数传参是形参，不会改变原变量的值。

	var a float32 = 2.12

	fmt.Printf("%0.4f\n", a)
}
