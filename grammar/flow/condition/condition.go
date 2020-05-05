package main

import (
	"fmt"
	"io/ioutil"
)

// GetFileContents 获取测试文件
func GetFileContents() ([]byte, error) {
	const fileName = "D:/go/src/atal.github.com/atal/learngo/grammar/flow/testfile.txt"
	return ioutil.ReadFile(fileName)
}

func ifCondition() {
	// if 条件语句中可以赋值，变量作用域只在 if 语句块中
	if contents, err := GetFileContents(); err != nil {
		fmt.Println("getFileContents err", err)
	} else {
		fmt.Printf("file contents:\n%s", contents)
	}
}

// switch 中每个 case 默认进行 break，如果不需要 break，需要写 fallthrough
func switchCondition() {
	const con = "x"
	var result int
	switch con {
	case "x":
		result = 1
	case "y":
		result = 2
	default:
		result = 0
	}
	fmt.Printf("switch result %v", result)
}

func main() {
	ifCondition()
	switchCondition()
}
