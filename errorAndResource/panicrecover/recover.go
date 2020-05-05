package main

import (
	"fmt"
	"log"
)

func recoverDemo() {
	// 定义一个匿名函数并且使用 defer 修饰调用
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			log.Printf(fmt.Sprintf("Error occurred and ignore:%v", err))
		} else {
			log.Printf(fmt.Sprintf("could not handle, repanic:%v", err))
			panic(r)
		}
	}()
	//panic(errors.New("real error"))
	panic("fake error")
}

func main() {
	recoverDemo()
}
