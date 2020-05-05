package deferdemo

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AtalH/learngo/functional/implinterface"
)

// deferStack 演示 defer 先进后出调用栈
func deferStack() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

// deferResource 演示 defer 在资源使用方面的应用
func deferResource() {
	const testFile = "D:\\go\\src\\github.com\\AtalH\\learngo\\errorAndResource\\fibfile.txt"
	file, err := os.Create(testFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// bufio 只是写到内存中，Flush 之后才会到硬盘
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	fib := implinterface.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Fprintln(writer, fib())
	}
}

// Test 运行测试
func Test() {
	deferStack()
	deferResource()
}
