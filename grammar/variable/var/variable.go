package main

import "fmt"

/*
方法外不能使用这种方式定义变量
v0 := false
*/

// 批量定义变量
var (
	vv1 = 1
	vv2 = "hello"
)

func main() {
	// 定义一个整型变量并赋值
	var v1 int = 1
	// 定义一个字符串变量并赋值
	var v2 string = "2"
	// 定义一个类型推断的变量
	var v3 = true
	// 省略 var 关键字的定义变量，具有类型推断
	v4 := 'c'
	fmt.Printf("v1=%d, v2=%s, v3=%t, v4=%c", v1, v2, v3, v4)

	fmt.Println()

	// 多个变量一行定义并赋值
	v5, v6 := 21, 3.141
	// 可使用 %v 格式化输出所有类型，使用 f 输出浮点型，.2 表示保留2位小数
	fmt.Printf("v5=%v, v6=%.2f\n", v5, v6)

	// 使用 %T 打印变量类型
	fmt.Printf("variable type of v4 is %T\n", v4)
	// 虽然 v6 已经定义过了，但是 v7 还没有定义，此处仍然可以使用 :=
	v6, v7 := 6, 7
	fmt.Printf("v6=%v, v7=%v\n", v6, v7)
}
