package main

import (
	"fmt"
)

func usage(chapterId int, childId int, chapterDesc string) {
	fmt.Println("----------------------------------------")
	fmt.Println("----- " + string(chapterId) + "_" + string(childId) + " ----- " + chapterDesc)
}

func notice(desc string) {
	fmt.Println("----------------------------------------")
	fmt.Println("----- 注意点：" + desc)
}

func main() {
	// chapter_1.1： go 变量
	usage(1, 1, "变量的声明，初始化，赋值，匿名变量用法")
	// go 创建变量的三种方式
	var x1 int = 10
	var x2 = 11
	x3 := 12 // := 运算符表示同时声明变量和初始化
	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
	notice("变量声明不需要分号作为语句结束符")
	notice("如果我们使用 := 声明变量，也不能重复定义已存在的同名变量，否则是一种编译错误")

	// go 可以多重赋值，例
	x1, x2 = x2, x1
	fmt.Println(x1)
	fmt.Println(x2)

	// 匿名变量
	usage(1, 1, "go 传参可以用 _ 表示匿名变量，用法灵活")
	var nickName string
	_, _, nickName = getName()
	fmt.Println(nickName)

	// 常量用法
	const Pi float64 = 3.14159265358979323846
	const zero = 0.0 // 无类型浮点常量
	const (
		size int64 = 1024
		eof        = -1 // 无类型整型常量
	)
	const u, v float32 = 0, 3   // u = 0.0, v = 3.0
	const a, b, c = 3, 4, "foo" // 多重赋值
	notice("go 中常量定义可以去限定类型，但不是必须的。如果没有指定，它就是无类型常量，换言之 2 可以是 int，int64, float 等多种类型")
	const mask = 1 << 10
	notice("常量的右值可以是一个编译期才能拿到的表达式，但不能是在运行期才能拿到的值，比如 const var1 = os.getEnv(\"Home\") 就会触发编译错误")
	fmt.Println(mask)

	usage(1, 1, "可以用 const 定义枚举量")
	const (
		Sunday = iota // itoa，true，false 三个为保留常量。itoa 初始为0，会在每次出现时 + 1, 在新的 const 定义出现时又会被重置为 0
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Saturday)

}

func getName() (firstName string, lastName string, nickName string) {
	return "Bill", "Gates", "aaa"
}
