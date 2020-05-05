package goroutine

import (
	"fmt"
	"time"
)

// Demo 演示 go 中协程并发的使用
func Demo() {
	for i := 0; i < 100; i++ {
		go func(index int) {
			for {
				fmt.Printf("printing from goroutine index %d\n", index)
			}
		}(i)
	}
	time.Sleep(time.Second)
}
