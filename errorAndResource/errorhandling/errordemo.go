package errorhandling

import (
	"fmt"
	"os"
)

// openFileDemo 根据 OpenFile 的定义，路径错误时，返回的 Error 是 PathError
func openFileDemo(fileName string) {
	// EXCL 和 CREATE 一起使用确保需要创建的文件之前是不存在的
	file, err := os.OpenFile(fileName, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			// 如果不是 pathError，无法继续处理就 panic
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()
}
