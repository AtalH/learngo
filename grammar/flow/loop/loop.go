package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func loop() {
	var sum int
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("loop result %v\n", sum)
}

// 此处 for 语句中省略了初始化部分，也可以只留中间的条件部分，甚至全都不需要，相当于 while
func convert2bin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func scanFile() {
	const fileName = "D:/go/src/atal.github.com/atal/learngo/grammar/flow1/testfile.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("read file error", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("file content line: %v\n", scanner.Text())
	}
}

func loopForever() {
	for {
		fmt.Println("looping...\n")
	}
}

func main() {
	loop()
	fmt.Printf("convert2bin result %v\n", convert2bin(13))
	scanFile()
}
