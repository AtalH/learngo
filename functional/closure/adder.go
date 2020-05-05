package closure

import "fmt"

// adder 这里 sum 是闭包中的自由变量，v 是局部变量
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

// TestAdder 测试闭包函数
func TestAdder() {
	a := adder()
	// 这个循环中不断调用 a(i)，内部的自由变量 sum 都不会被重置为 0
	for i := 1; i < 10; i++ {
		result := a(i)
		fmt.Printf("0 + ... + %d = %d\n", i, result)
	}
}
