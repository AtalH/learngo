package implinterface

import "fmt"

// Fibonacci 使用函数闭包实现斐波那契数列
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// TestFib 测试
func TestFib() {
	fib := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}
