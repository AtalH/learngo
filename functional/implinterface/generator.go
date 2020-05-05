package implinterface

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// intGenerator 是一个 func() int 函数类型
type intGenerator func() int

// Read 为函数类型 intGenerator 实现 io.Read 接口
func (g intGenerator) Read(p []byte) (n int, err error) {
	// g 是一个 func() int 函数类型
	next := g()
	// 转为字符串
	str := fmt.Sprintf("%d\n", next)
	// 用 strings.NewReader 代理，把 str 放到 p 字节数组中
	// TODO: 此处 p 的容量可能不够
	return strings.NewReader(str).Read(p)
}

func readLines(reader io.Reader, maxLine int) {
	scanner := bufio.NewScanner(reader)
	for i := 0; i < maxLine; i++ {
		scanner.Scan()
		fmt.Println(scanner.Text())
	}
}

// fib 返回的函数是 intGenerator 类型
func fib() intGenerator {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// TestGenerator 测试生成器
func TestGenerator() {
	fmt.Println("test int generator...")
	// fib() 返回值是 intGenerator 类型的函数
	f := fib()
	readLines(f, 10)
}
