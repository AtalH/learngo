package main

import (
	"fmt"
	"math"
)

// 方法外定义常量，未指定类型时，类型不确定
const c = "con"

func constFunc() {
	i := 3
	// 错误，math.Sqrt 接收的是 float64 类型
	//j := math.Sqrt(i)

	const (
		n        = 3
		m string = "String"
	)
	// 由于 n 定义时未指定类型，此时当作 float64 类型处理
	j := math.Sqrt(n)
	fmt.Printf("j result %v\n", j)
	fmt.Printf("%s print %d string %v\n", c, i, m)
}

/*
iota 是一个常量定义中使用的自增变量，从 0 开始，每换一行自增 1，遇到 const 关键字就重置为 0。
常量定义时，后面的行如果省略赋值，会跟前一行一样的赋值（但是 iota 会自增）
*/
func enums() {
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Printf("b %d\n", b)
	fmt.Printf("kb %d\n", kb)
	fmt.Printf("mb %d\n", mb)
	fmt.Printf("gb %d\n", gb)
	fmt.Printf("tb %d\n", tb)
	fmt.Printf("pb %d\n", pb)

	/*
		遇到 const 关键字，iota 重置为 0
		_ = iota 下划线可以接收值，但不使用，可以用在多返回值的函数调用，只希望使用一个返回值时
	*/
	const (
		i0 = iota
		i1
		_
		i3
	)
	fmt.Printf("iota i0 = %d, i1 = %d, i3 = %d\n", i0, i1, i3)
}

func main() {
	constFunc()
	enums()
}
