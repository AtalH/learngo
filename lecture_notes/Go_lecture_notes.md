# Go 学习课程笔记

## useful link

- https://golang.google.cn/
- learn https://www.liwenzhou.com/
## 安装

- Windows 下 msi 安装包会把程序安装在  C:\Go
- 配置自定义环境变量
    - GOROOT 指 go 语言的安装路径，也就是 C:\Go，确保系统环境变量 PATH 中有 %GOROOT%\bin
    - GOPATH  是用户代码存放的目录，可以自定义，确保用户环境变量中有 GOPATH，并且 PATH 中有 %GOPATH%\bin，这样编译 go 代码后，就可以在 cmd 中直接执行编译后的文件。
- GOPATH 下的代码目录
    - 代码总是会保存在$GOPATH/src目录下。在工程经过go build、go install或go get等指令后，会将下载的第三方包源代码文件放在$GOPATH/src目录下
    - 编译生成的二进制可执行文件放在 $GOPATH/bin目录下
    - 生成的中间缓存文件会被保存在 $GOPATH/pkg 下
- 项目结构
    - $GOPATH/src/域名/用户/项目/模块
    - $GOPATH/src/域名/部门或组/用户/项目/模块
- go mod
- go proxy
    - https://goproxy.io/zh/
    - 对于 go 版本在 1.13 以上的，如下设置完下面几个环境变量后，go 命令将从公共代理镜像中快速拉取所需的依赖代码。
    ```
    go env -w GO111MODULE=on
    go env -w GOPROXY=https://goproxy.io,direct
    
    # 设置不走 proxy 的私有仓库，多个用逗号相隔（可选）
    go env -w GOPRIVATE=*.corp.example.com
    ```
- gopm
    - go package manager
    - 有镜像功能，能够下载 golang.org 的包
    - doc https://github.com/gpmgo/docs/tree/master/zh-CN
    - 安装
        ```
        go get -v -u github.com/gpmgo/gopm
        ```
    - 需要安装 git
    - 使用
        ```
        gopm get -h
        gopm get -g -u -v golang.org/xxxx
        ```
- 环境变量配置错误
    - 在 go 版本 1.13，如果 go env -w GO111MODULE=ON，识别为错误的环境变量，但又没有删除这个环境变量的命令，只能手工删除。
    - Windows 系统中，打开 Users\\<User Name>\AppData\Roaming\go\env 文件修改
## IDE

- vs code 配置
    - 在 vs code 中 ctrl+shift+p，输入 go:install，选择 Go:Install/Update Tools 命令，勾选所有插件，回车安装。
    - 如果配置了 go proxy，则能正常安装成功。
    - 如果不想使用 go proxy，则可以如下下载插件源码，再进行安装
    ```
    创建 GOPATH/src/golang.org/x 目录
    git clone https://github.com/golang/tools.git tools
    git clone https://github.com/golang/lint.git
    ```
## 入门

- 基本概念
    - 包 package 是在同一个文件夹下的代码集合，同一个包下的函数、类型、变量、常量都是互相可见的
    - 模块 module 是一起发布的、有相关性的package
    - 仓库 repository 是多个 module 的集合
    - module 的路径不仅是 package的 import 路径的前缀，而且也指示了 go 命令应该去哪里下载
        - 比如模块 golang.org/x/tools ，应该去 https://golang.org/x/tools 这个 repository 下载
    - package 的 import 路径是 module 路径加上 module 内的一个子目录
    - 标准库的 import 路径没有 module 前缀
    - 一个 package 下只有一个 main
### 基础语法
- 变量
    - 变量定义
        - 变量名在前，类型在后
        - 相同类型的变量可一行定义赋值
        ```
        var a, b int = 1， 2
        ```
        - 有类型推导，省略类型
        ```
        var a, b = 1, 1.2
        ```
        - 方法内本地变量可省略 var
        ```
        a, b , v6:= 1, 1.2, 6
        // 虽然 v6 已经定义过了，但是 v7 还没有定义，此处仍然可以使用 :=
    	v6, v7 := 6, 7
    	fmt.Printf("v6=%v, v7=%v\n", v6, v7)
        ```
        - 可批量定义
        ```
        var (
            a = 1
            b = 2.1
        )
        ```
    - fmt 打印
        - fmt.Printf() 格式化打印
            - %s 格式化字符串
            - %d 格式化整型
            - %f 格式化浮点型
            - %c 格式化字符型
            - %t 格式化布尔型
            - %T 格式化变量类型
            - %v 自动推断类型
        ```
        fmt.Printf("v1=%d, v2=%s, v3=%t, v4=%c", v1, v2, v3, v4)
        fmt.Println()
        fmt.Printf("v5=%v, v6=%.2f\n", v5, v6)
        fmt.Printf("variable type of v4 is %T", v4)
        ```
        - fmt.Fprint 是写到 io.Writers，而 fmt.Printf 是写到 os.Stdout
        - fmt.Sprint 是返回格式化的字符串
- 内建类型
    - bool
    - string
        - 字符块用反点
        ``` go
        var strBlock = `abc"qqq"
                        efd`
        ```
    - int, uint, int8, uint8, ... int64, uint64
        
        - 整型，uint 表示无符号整型，长度由操作系统决定，32 位系统中就是 32 位的；int8 指有符号 8 位长度的整型
    - uintptr
        
        - 指针
    - byte
    - rune
        
        - 32 位（4 字节）字符型，单字节 char 型对 utf8 兼容不好
    - float32, float64
        
        - 没有 double 类型
    - 浮点数都有不精确的问题(IEEE754 浮点数表示标准)，java 中是使用 BigDecimal 类解决浮点数精确计算和大数运算问题，go 中使用 math/big 解决
    ```
    func bigNum() {
    	f1 := 0.09
    	f2 := 2.1
    	// 结果是 0.189
    	fmt.Printf("f1*f2 result %v\n", f1*f2)
    	// 强制类型转换后得0
    	var i int = int(f1 * f2)
    	fmt.Printf("i result %v\n", i)
    
    	// 使用 math/big 可以精确计算
    	df1 := big.NewFloat(f1)
    	df2 := big.NewFloat(f2)
    	res := big.NewFloat(0)
    	// 将计算结果存入 res 中，覆盖原值
    	res.Mul(df1, df2)
    	fmt.Printf("df result %v", res)
    }
    ```
    - complex64, complex128
        - 复数类型 
        ```
        不能写做 4*i，否则 i 会识别成变量
        var c = 3 + 4i
        ```
        - 只有 i 时，写成 1i
        ```math
        1i^2 = -1
        ```
        - 欧拉公式
        ```math
        e^{i\pi} + 1=0
        ```
    - 类型转换
        - go 中只有强制、显式的类型转换，没有隐式类型转换
        ```
        var i int= int(0.09f * 2.1f) 
        ```
- 常量
    
    - 常量定义时不指定类型的话，可作为各种类型使用
```
package main

import (
	"fmt"
	"math"
)

// 方法外定义常量，未指定类型时，类型不确定
const c = "con"

func main() {
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
```
- 枚举
    - 使用常量来实现枚举
    ```
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
    		_ = iota 下划线可以接收值，但不使用，也可以用在多返回值的函数调用，只希望使用一个返回值时
    	*/
    	const (
    		i0 = iota
    		i1
    		_
    		i3
    	)
    	fmt.Printf("iota i0 = %d, i1 = %d, i3 = %d\n", i0, i1, i3)
    }
    ```
- 条件
    - if
    ```
    func getFileContents() ([]byte, error) {
    	const fileName = "D:/go/src/atal.github.com/atal/learngo/grammar/flow/testfile.txt"
    	return ioutil.ReadFile(fileName)
    }
    
    func ifCondition() {
    	// if 条件语句中可以赋值，变量只在 if 语句块中有效
    	if contents, err := getFileContents(); err != nil {
    		fmt.Println("getFileContents err", err)
    	} else {
    		fmt.Printf("file contents:\n%s", contents)
    	}
    }
    ```
    - switch
        - switch 后也可以不跟变量，直接 case 判断条件
    ```
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
    ```
- 循环
    ```
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
    
    // go 中没有 while
    func loopForever() {
    	for {
    		fmt.Println("looping...")
    	}
    }
    ```
- 函数
    - 函数一等公民，可作为函数参数传递，甚至将匿名函数当参数
    - 支持多返回值
    - error 当作返回值进行传递，没有 try...catch 语法
    - 有可变参数列表
    - 无默认参数，无可选参数语法，无重载
        ```
        func plug(a , b int) int {
            return a + b
        }
        ```
- 指针
    ```
    // 定义一个 int 型变量 a
    var a int
    // 定义整型指针类型变量 pa，并赋值为变量 a 的地址
    var pa *int = &a
    ```
    - go 中都是值传递，Java 中除了基本类型，都是引用传递。
    - go 可以通过传递指针实现引用传递效果

## 内置容器

- 数组
    - 定义数组
        ```
        // 定义 5 个元素的 int 型数组，初始值是 0
        var arr1 [5]int
        
        // 定义一个自动推断长度的数组
        var arr2 := [...]int{0, 1, 2, 3}
        
        // 定义一个 2 维数组，4 行 5 列
        var arr3 [4][5] int
        ```
    - go 中数组是值类型，函数传递时是值传递，因此会进行值拷贝，对大数组需要注意此问题。而 java 中是引用传递
        ```
        // 定义一个接收长度为 5 的 int 型数组
        func f1(arr [5]int){
            
        }
        ```
    - 通过传递数组指针，也可以实现引用传递
    - 数组的遍历
        - 通过 for 索引遍历
            ```
            func forArray(arr [4]int) {
            	for i := 0; i < len(arr); i++ {
            		fmt.Printf("arr index %d value %d\n", i, arr[i])
            	}
            }
            ```
        - 通过 range 遍历
            ```
            func rangeArray(arr [5]int) {
            	for i, v := range arr {
            		fmt.Printf("arr index %d value %d\n", i, v)
            	}
            }
            ```
- 切片
    - 切片是数组的一个视图，对切片的修改，会影响源数组。因此通过函数传递切片，是引用传递
        ```
        // 数组
        arr : = [...]int{0, 1, 2, 3, 4, 5, 6}
        
        // 定义一个切片，并赋值为数组 arr 的第 2 到 第 3 个元素，也就是前闭后开区间
        arrSlice := arr[2:4]
        
        // 定义一个 int 型 slice
        var arrSlice1 []int = arr[:3]
        
        // 可以对切片再次切片
        arrSlice1 := arrSlice1[:]
        ```
    - 切片的扩展。切片内部维护了一个指向头部的指针 ptr、一个长度值 len，一个容量值 cap
        - 对切片进行重新切片时，下标可以向后超出自身的长度值 len，但不能超出容量值 cap
        ```
         var arr [...]int := {0, 1, 2, 3, 4, 5, 6}
         
         // slice 值为 {1, 2, 3}
         slice := arr[1:4]
         
         // 索引 6 超出了变量 slice 的长度值 len，但没有超出 arr 数组的容量值 cap
         // slice1 值为 {2, 3, 4}
         slice1 := slice[1:4]
        ```
    - 切片的 append
        ```
        var arr = [...]int{0, 1, 2, 3, 4, 5, 6}

        func sliceAppend() {
        	slice := arr[3:5]
        	// 对 slice 进行追加，会改变源数组的内容
        	slice0 := append(slice, 50)
        	slice1 := append(slice0, 60)
        	// 当追加超出源数组的长度时，底层会给这个 slice 创建一个新的数组
        	// 扩容 cap 按 2 的指数增加
        	slice2 := append(slice1, 70)
        	fmt.Printf("\nafter append, arr=%v\n", arr)
        	fmt.Println(slice, "\n", slice0, "\n", slice1, "\n", slice2)
        }
        ```
    - 创建切片
        ```
        func createSlice() {
    	fmt.Println("in createSlice()")
    	// 定义一个 int 型 slice ，零值是 nil
    	var slice []int
    	// nil 也不会错误
    	slice = append(slice, 1)
    
    	// 定义一个 int 型 slice
    	slice0 := []int{2, 4, 5, 6}
    
    	// 创建一个空的 slice，并指定长度为5（容量也是5）
    	slice1 := make([]int, 5)
    
    	// 创建一个空的 slice，并指定长度为5，容量是10
    	slice2 := make([]int, 5, 10)
    	}
    	```
    - 切片的删除
        ```
        func sliceDel() {
        	fmt.Println("in sliceDel()")
        	slice := arr[:]
        	// 删除索引为 3 的元素
        	slice = append(slice[:3], slice[4:]...)
        	fmt.Printf("slice value=%v, len=%d, cap=%d\n", slice, len(slice), cap(slice))
        
        	// 删除第一个元素
        	slice = slice[1:]
        	fmt.Printf("slice value=%v, len=%d, cap=%d\n", slice, len(slice), cap(slice))
        
        	// 删除最后一个元素
        	slice = slice[:len(slice)-1]
        	fmt.Printf("slice value=%v, len=%d, cap=%d\n", slice, len(slice), cap(slice))
        }
        ```
    - 切片的复制
        ```
        func sliceCopy() {
        	// 此处只是定义了一个为 nil 的 slice，没有空间，无法拷贝
        	//var targetSlice []int
        	// 使用 make 得到 slice 容量为 7
        	targetSlice := make([]int, 7)
        	srcSlice := arr[:]
        	copy(targetSlice, srcSlice)
        	fmt.Printf("srcSlice=%v, targetSlice=%v\n", srcSlice, targetSlice)
        }
        ```
- map
    - 定义 map
        ```
        // 定义一个 key 为 string 类型，value 为 int 类型的 map 并赋值
        var map0 = map[string]int{
        	"k0": 0,
        	"k1": 1,
        }
        
        func defMap() {
        	// 创建一个为空 map
        	map1 := make(map[int]string)
        
        	//定义一个 map，零值为 nil
        	var map2 map[string]int
        	fmt.Printf("map0=%v, map1=%v, map2=%v\n", map0, map1, map2)
        }
        ```
    - map 的操作
        ```
        func mapOp() {
        	// 取得 key 为 k0 的 value
        	v0 := map0["k0"]
        	// 不存在的 key，返回值为 value 类型 int 零值 0
        	v3 := map0["noKey"]
        
        	// 可以返回两个值
        	if v, isExist := map0["k1"]; isExist {
        		fmt.Printf("key k1 exist, value is %v\n", v)
        	} else {
        		fmt.Printf("key k1 dose not exist")
        	}
        
        	delete(map0, "k1")
        	// 删除不存在的 key 不会报错
        	delete(map0, "noKey")
        
        	fmt.Printf("v0=%v, v3=%v\n", v0, v3)
        	fmt.Printf("map0=%v\n", map0)
        
        	// 增加 key value
        	map0["k3"] = 3
        	fmt.Printf("map0 len=%d", len(map0))
        	// map 的遍历，不保证顺序
        	for k, v := range map0 {
        		fmt.Printf("map0 has key=%v, value=%v\n", k, v)
        	}
        }
        ```
    - map 对 key 的要求
        - 低层使用哈希表，key 必须能比较相等
        - 除了 slice、function、map 类型，其他内建类型都可以做 key
        - 当 Struct 不包含 slice、function、map 类型时，可做 key
- 字符串与字符 rune
    ```
    var chStr = "hi中国啊。"

    func strContent() {
    	// 这里输出的是字节数 14
    	fmt.Printf("str len=%d\n", len(chStr))
    	fmt.Println("十六进制字节")
    	// 中文字符占 3 个字节
    	for _, b := range []byte(chStr) {
    		fmt.Printf("%X ", b)
    	}
    	fmt.Println()
    
    	// 按 rune 字符输出 Unicode 十六进制，索引仍然是字节索引
    	fmt.Println("rune X Unicode format")
    	for i, s := range chStr {
    		fmt.Printf("index=%d, rune=%X \n", i, s)
    	}
    	fmt.Println()
    
    	// 按 rune 字符输出 utf8 解码，但索引仍然是字节索引
    	fmt.Println("rune format")
    	for i, s := range chStr {
    		fmt.Printf("index=%d, rune=%c \n", i, s)
    	}
    }
    
    func runeUtf8() {
    	strBytes := []byte(chStr)
    	fmt.Printf("rune count=%d", utf8.RuneCount(strBytes))
    	for len(strBytes) > 0 {
    		// 从 strBytes 中按 utf8解码第一个字符，size 是字符的字节数
    		ch, size := utf8.DecodeRune(strBytes)
    		fmt.Printf("utf8 decode a rune=%c, rune size=%d\n", ch, size)
    		strBytes = strBytes[size:]
    	}
    
    	// 最直接的按 rune 进行索引，转 rune 数组后，每一个元素都按 4 字节存了
    	for i, ch := range []rune(chStr) {
    		fmt.Printf("%d=%c ", i, ch)
    	}
    }
    ```
    - 字符串工具包
        - strings 包括 Split、join、Contains、Trim 等等工具方法
## 结构体

- 使用结构体能实现部分面向对象
- 结构体只支持封装，不支持继承和多态，而通过接口进行扩展
    ```
    // 定义一个结构体
    type treeNode struct {
    	value       string
    	left, right *treeNode
    }
    
    // 结构体没有构造函数，可以通过定义工厂方法来创建
    // 虽然返回的是方法内部的变量地址，但是不会像 C++ 那样被回收掉
    func createTreeNode(value string) *treeNode {
    	return &treeNode{value: value}
    }
    
    // 各种实例化结构体方法
    func buildATree() *treeNode {
    	var rootNode treeNode
    	rootNode = treeNode{value: "Root"}
    	rootNode.left = &treeNode{}
    	rootNode.right = &treeNode{"root's right node", nil, nil}
    	rootNode.right.left = new(treeNode)
    	rootNode.left.right = createTreeNode("a right node 2")
    	return &rootNode
    }
    
    // 给结构体 treeNode 定义一个方法
    // 与 func print(node treeNode) 定义方法的区别仅仅是调用形式不同
    // 仍然是值传递，node 在方法内部不会是 nil
    func (node treeNode) print() {
    	fmt.Println("node's value = ", node.value)
    }
    
    // 通过传递指针类型，实现引用传递
    func (node *treeNode) setValue(value string) {
    	if node == nil {
    		return
    	}
    	// 尝试给 value 赋值，必须先检查 nil 不然会报错
    	node.value = value
    }
    
    // 中序遍历
    func (node *treeNode) traverse() {
    	if node == nil {
    		return
    	}
    	node.left.traverse()
    	node.print()
    	node.right.traverse()
    }
    
    func main() {
    	tree := buildATree()
    	// setValue 是引用传递，print 是值传递，调用方式一样
    	tree.left.setValue("root's left node")
    	tree.left.print()
    
    	// nil 变量也可以调用方法而不会报错
    	var aNilNode treeNode
    	// 如果 print 方法也是引用传递，取 value 值也不会报错
    	aNilNode.print()
    	// setValue 方法尝试给 value 赋值，必须先检查 nil 不然会报错
    	aNilNode.setValue("nil node set value")
    
    	fmt.Println("中序遍历开始...")
    	tree.traverse()
    }
    ```
#### 包与权限
- 同一个目录下的所有文件的 package 都一样
- 结构名、变量名、方法名首字母大写时，对其他 package 可见，即 public
- 结构体的方法必须定义在同一个 package 内
    - public 的结构体和方法，必须要有注释，并且注释要以结构体名或方法名开头
    - package 的注释要以 "Package + 包名" 开头
    - 文件 tree/node.go
        ```
        // Package tree 树结构相关包
        package tree

        import "fmt"
        
        // Node 树节点
        type Node struct {
        	Value       int
        	Left, Right *Node
        }
        
        // Print 打印 value 值
        func (node *Node) Print() {
        	fmt.Println(node.Value)
        }
        ```
    - 文件 tree/traversal.go
        ```
        package tree

        // Traverse 中序遍历打印
        func (node *Node) Traverse() {
        	if node == nil {
        		return
        	}
        	node.Left.Traverse()
        	node.Print()
        	node.Right.Traverse()
        }
        ```
    - 使用结构体 Node 时，必须用包名 tree 开头引用
        ```
        package main

        import "atal.github.com/atal/learngo/package/tree"
        
        func main() {
        	node := tree.Node{Value: 0}
        	node.Print()
        }
        ```
#### 扩展已有类型
- 可以多系统类型或第三方类型进行扩充
- 方法一，使用组合方式进行扩展
    ```
    // Xnode 对 tree.Node 进行扩展
    type Xnode struct {
    	Node *tree.Node
    }
    
    // PostOrder 扩充 tree.Node 添加后续遍历方法
    func (myTreeNode *Xnode) PostOrder() {
    	if myTreeNode == nil || myTreeNode.Node == nil {
    		return
    	}
    	//将 node 类型的 Left 转换为 Xnode 类型，
    	//必须定义变量，不能直接调用
    	left := Xnode{myTreeNode.Node.Left}
    	right := Xnode{myTreeNode.Node.Right}
    
    	left.PostOrder()
    	right.PostOrder()
    	myTreeNode.Node.Print()
    }
    ```
- 使用别名方式进行扩展
    ```
    package queue

    // IntQueue 使用别名方式扩充slice实现整型队列
    type IntQueue []int
    
    // Push 入队
    func (q *IntQueue) Push(v int) {
    	*q = append(*q, v)
    }
    
    // Pop 出队
    func (q *IntQueue) Pop() int {
    	tail := (*q)[0]
    	*q = (*q)[1:]
    	return tail
    }
    
    // IsEmpty 队列是否为空
    func (q *IntQueue) IsEmpty() bool {
    	return len(*q) <= 0
    }
    
    func main() {
    	q := queue.IntQueue{0}
    	q.Push(1)
    	q.Push(2)
    	fmt.Println("is q empty:", q.IsEmpty())
    	fmt.Println("q pop:", q.Pop())
    	fmt.Println("q pop:", q.Pop())
    	fmt.Println("is q empty:", q.IsEmpty())
    	fmt.Println("q pop:", q.Pop())
    	fmt.Println("is q empty:", q.IsEmpty())
    }
    ```
## 接口

- duck typing
- 定义接口
    ```
    package face

    // Retriever 定义了一个包含 Get 方法的接口
    type Retriever interface {
    	Get(url string) string
    }
    ```
- 实现接口
    - 实现一个接口，只需要实现接口具有的方法，其他方面跟接口没有关系了，不像 java 那样强制实现
    ```
    package qq

    import (
    	"net/http"
    	"net/http/httputil"
    	"time"
    )
    
    // Retriever qq 包中的 Retriever 实现类
    type Retriever struct {
    	UserAgent string
    	TimeOut   time.Duration
    }
    
    // Get 实现 Get 方法
    func (r Retriever) Get(url string) string {
    	resp, err := http.Get(url)
    	if err != nil {
    		panic(err)
    	}
    	body, err := httputil.DumpResponse(resp, true)
    	if err != nil {
    		panic(err)
    	}
    	resp.Body.Close()
    	return string(body)
    }
    ```
- 接口的使用
    ```
    package main

    import (
    	"fmt"
    
    	"atal.github.com/atal/learngo/inter/face"
    	"atal.github.com/atal/learngo/inter/qq"
    )
    
    // getQQ 虽然参数是 face.Retriever，但是只要是有 Get 方法的都可以传进去
    func getQQ(r face.Retriever) string {
    	return r.Get("https://qq.com/")
    }
    
    func main() {
    	var retriever face.Retriever
    	retriever = qq.Retriever{}
    	page := getQQ(retriever)
    
    	//或简单的使用方法
    	page = getQQ(qq.Retriever{})
    	fmt.Println(page)
    }
    ```
- 接口的值类型
    ```go
    func showInterface() {
    	var retriever face.Retriever
    	retriever = qq.Retriever{UserAgent: "Chrome", TimeOut: 10}
    	fmt.Printf("type:%T, value:%v\n", retriever, retriever)
    	// 当接口方法是指针接收者是，type就是指针类型，value就是地址类型
    }
    ```
    - 通过 switch 语句判断
    ```go
    func typeSwith() {
    	var retriever face.Retriever
    	retriever = qq.Retriever{UserAgent: "FireFox", TimeOut: 10}
    	switch v := retriever.(type) {
    	case qq.Retriever:
    		fmt.Println("qq.Retriever contents:", v.UserAgent)
    	case face.Retriever:
    		fmt.Println("face.Retriever")
    	}
    }
    ```
    - 通过 typeAssertion
    ```go
    func typeAssertion() {
    	var retriever face.Retriever
    	retriever = qq.Retriever{UserAgent: "FireFox", TimeOut: 10}
    	if r, ok := retriever.(qq.Retriever); ok {
    		fmt.Println("this is a qq.Retriever, UserAgent:", r.UserAgent)
    	} else {
    		fmt.Printf("this is not a qq.Retriever, %T\n", r)
    	}
    }
    ```
    - 接口变量自带指针。接口变量同样采用值传递，几乎不需要使用接口的指针
    - 定义一个支持任何类型的 Queue，interface{}表示任何类型
    ```go
    package inter

        // Queue interface{}表示任何类型
        type Queue []interface{}
        
        // Push 入队
        func (q *IntQueue) Push(v interface{}) {
        	*q = append(*q, v)
        }
        
        // Pop 出队
        func (q *IntQueue) Pop() interface{} {
        	tail := (*q)[0]
        	*q = (*q)[1:]
        	return tail
        }
    ```
- 接口的组合
    - 如下实现者 HTTPUtil 也不需要指定实现了什么接口（HTTPClient）
    ```go
    package compose

    // Getter 定义 get 方法
    type Getter interface {
    	Get(url string) string
    }
    
    // Poster 定义 post 方法
    type Poster interface {
    	Post(url string) string
    }
    
    // HTTPClient 组合了 Get() Post() 方法，也可以定义自己的方法
    type HTTPClient interface {
    	Getter
    	Poster
    	Session(key string, value string) bool
    }
    
    // HTTPUtil 是 HTTPClient 接口的实现类
    type HTTPUtil struct {
    	url string
    }
    
    // Get 是 HTTPUtil 中的具体方法实现
    func (http HTTPUtil) Get(url string) string {
    	return "fake get result"
    }
    
    // Post 是 HTTPUtil 中的具体方法实现
    func (http HTTPUtil) Post(url string) string {
    	return "fake post result"
    }
    
    // Session 是 HTTPUtil 中的具体方法实现
    func (http HTTPUtil) Session(key string, value string) bool {
    	return true
    }
    ```
    - io.ReadWriteCloser 就是组合了 Read Write Closer 接口
- 常用系统接口
    - Stringer 接口中的 String() 方法相当于 Java 中的 同String()
    ```go
    package sysface

    import "fmt"
    
    // Str 实现了 Stringer 接口
    type Str struct {
    	Content string
    }
    
    // String 实现 Stringer 接口的 String() 方法
    func (s Str) String() string {
    	return fmt.Sprintf("str's content is [%v]", s.Content)
    }
    ```
    - io.Reader
    - io.Writer
## 函数式编程

- 闭包
    - 如下 adder() 方法返回值是一个函数 func(int) int，这个函数有一个 int 类型参数，返回值是 int 类型。
    - sum 变量属于 adder() 作用域，但是由于闭包的存在，使得 func(v int) int 中对 sum 的修改都能够保存下来
    - 此处 sum 称作自由变量，v 称作局部变量
    - 在 TestAdder() 方法的 for 循环中，不断调用 a(i)，内部的自由变量 sum 都不会被重置为 0，而是有状态的不断累加
	- 此处可类比 java，把 a 看作 adder 类型实例，sum 就是内部成员变量，只要实例 a 不被销毁，sum 状态就不会被重置
	- 如果还有一个 b := adder()，那么 b 的 sum 和 a 的 sum 是互不影响的
	- 在“正统”函数式编程中，是不允许有自由变量存在的，只能是常量和函数
    ```go
    package closure

    import "fmt"
    
    // adder 这里 sum 是闭包中的自由变量，v 是局内部变量
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
    ```
    - python、JavaScript、C++11 都支持闭包，java8 之后可以通过 Function 接口和 lambda 实现闭包。java8 之前可以通过匿名函数实现闭包，java 中的自由变量不能用基本类型
- 应用案例1
    - 斐波那契数列
    ```go
    func fibonacci() func() int {
    	a, b := 0, 1
    	return func() int {
    		a, b = b, a+b
    		return a
    	}
    }
    
    // TestFib 测试
    func TestFib() {
    	fib := fibonacci()
    	for i := 0; i < 10; i++ {
    		fmt.Println(fib())
    	}
    }
    ```
- 应用案例2
    - 为函数实现接口
    - 实现一个生成器
    ```go
    
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

    ```
- 应用案例 3
    - 通过传入函数，让调用者可以自定义遍历时的操作
    ```go
    // Node 树节点
    type Node struct {
    	Value       int
    	Left, Right *Node
    }
    
    // onTraverse 给结构体 Node 添加一个中序遍历函数，
    // 参数 f 是一个接受 Node 参数的函数
    func (node *Node) onTraverse(f func(*Node)) {
    	if node == nil {
    		return
    	}
    	node.Left.onTraverse(f)
    	f(node)
    	node.Right.onTraverse(f)
    }
    
    // TestOnTraverse 测试
    func TestOnTraverse() {
    	fmt.Println("testing TestOnTraverse")
    	left := Node{Value: 1}
    	right := Node{Value: 2}
    	root := Node{Value: 0, Left: &left, Right: &right}
		//调用者可以自定义遍历时是输出节点值还是做其他操作
    	root.onTraverse(func(n *Node) {
    		fmt.Println(n.Value)
    	})
    }
    ```
## 错误处理与资源管理

- defer
    - defer 能够确保其修饰的语句在函数结束时被调用
    - 一个函数中有多个语句被 defer 修饰时，按照栈先进后出的顺序执行
    - defer 语句中参数变量，会在当时确定下来，而不是到执行时才确定
    - 常用于资源的关闭处理，在开启资源后就使用 defer 修饰关闭资源的语句
    - 如下情况可以使用 defer
        - Open/Close
        - Lock/Unlock
        - PrintHeader/PrintFooter
    ```go
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
    ```
- error
    ```go
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
    ```
    - 文件服务器
        - 访问 http://127.0.0.1:8180/list/errorAndResource/webserver/listfile.go
        - http server 中发生 panic 不会停止程序
    ```go
    package main

    import (
    	"io/ioutil"
    	"log"
    	"net/http"
    	"os"
    )
    
    const pathOfListFilePrefix string = "/list/"
    
    // requestHanlder 定义一个拥有返回 error 的函数类型
    type requestHanlder func(writer http.ResponseWriter, request *http.Request) error
    
    // handlerWrapper 将 requestHanlder 包装，统一处理其错误信息，返回的函数符合 http.HandleFunc 要求
    func handlerWrapper(hanler requestHanlder) func(writer http.ResponseWriter, request *http.Request) {
    	return func(writer http.ResponseWriter, request *http.Request) {
    		err := hanler(writer, request)
    		if err == nil {
    			return
    		}
    		log.Printf("handle request [%s] error: %s", request.URL, err.Error())
    		statusCode := http.StatusOK
    		switch {
    		case os.IsNotExist(err):
    			statusCode = http.StatusNotFound
    		case os.IsPermission(err):
    			statusCode = http.StatusForbidden
    		default:
    			statusCode = http.StatusInternalServerError
    		}
    		http.Error(writer, http.StatusText(statusCode), statusCode)
    	}
    
    }
    
    func handleFileList(writer http.ResponseWriter, request *http.Request) error {
    	// 截取 ulr 后的路径作为文件的路径
    	filepath := request.URL.Path[len(pathOfListFilePrefix):]
    	file, err := os.Open(filepath)
    	if err != nil {
    		return err
    	}
    	defer file.Close()
    	contents, err := ioutil.ReadAll(file)
    	if err != nil {
    		return err
    	}
    	writer.Write(contents)
    	return nil
    }
    
    func main() {
    	// 设定一个 url 及其处理函数
    	http.HandleFunc(pathOfListFilePrefix, handlerWrapper(handleFileList))
    	// 开启端口
    	err := http.ListenAndServe(":8180", nil)
    	if err != nil {
    		panic(err)
    	}
    }
    ```
- panic and recover
    - panic 是个严重的错误，一般程序中不应出现
        - 意料之中的错误使用 error
    - 程序中遇到 panic
        - 停止当前程序执行
        - 一直向上返回，执行每一层的 defer
        - 如果没有遇见 recover，程序退出
    - recover 
        - 仅在 defer 修饰的语句中调用
        - 能够获取 panic 的值
        - 根据 panic 的值，如果无法处理，则可以重新 panic
    ```go
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
    ```
- 文件服务器的改进
    - 自定义错误类型
    ``` go
    // UserError 自定义接口类型：用户错误
    type UserError interface {
    	error
    	Message()
    }
    ```
    - 错误类型的实现
    ```go
    type userErr string

    // Error 实现 error 的接口函数
    func (e userErr) Error() string {
    	return e.Message()
    }
    
    // Message 实现 UserError 接口的函数
    func (e userErr) Message() string {
    	return string(e)
    }
    ```
    - 在意料之中错误中使用自定义错误类型
    ```go
    func handleFileList(writer http.ResponseWriter, request *http.Request) error {
    	// 截取 ulr 后的路径作为文件的路径
    	url := request.URL.Path
    	if strings.Index(url, pathOfListFilePrefix) != 0 {
    		return userErr("path should start with " + pathOfListFilePrefix)
    	}
    	filepath := request.URL.Path[len(pathOfListFilePrefix):]
    	file, err := os.Open(filepath)
    	if err != nil {
    		return err
    	}
    	defer file.Close()
    	contents, err := ioutil.ReadAll(file)
    	if err != nil {
    		return err
    	}
    	writer.Write(contents)
    	return nil
    }
    ```
    - 错误处理增加 recover 和自定义错误类型的处理
    ```go
    // handlerWrapper 将 requestHanlder 包装，统一处理其错误信息，返回的函数符合 http.HandleFunc 要求
    func handlerWrapper(hanler requestHanlder) func(writer http.ResponseWriter, request *http.Request) {
    	return func(writer http.ResponseWriter, request *http.Request) {
    		// recover 代码块，处理一些无法意料的错误
    		defer func() {
    			if r := recover(); r != nil {
    				log.Printf("panic: %v", r)
    				http.Error(writer, http.StatusText(http.StatusInternalServerError),
    					http.StatusInternalServerError)
    			}
    		}()
    		err := hanler(writer, request)
    		if err == nil {
    			return
    		}
    		log.Printf("handle request [%s] error: %s", request.URL, err.Error())
    		// 对于可以展示给用户的 userErr，直接返回
    		if ue, ok := err.(userErr); ok {
    			http.Error(writer, ue.Message(), http.StatusBadRequest)
    			return
    		}
    		statusCode := http.StatusOK
    		switch {
    		case os.IsNotExist(err):
    			statusCode = http.StatusNotFound
    		case os.IsPermission(err):
    			statusCode = http.StatusForbidden
    		default:
    			statusCode = http.StatusInternalServerError
    		}
    		http.Error(writer, http.StatusText(statusCode), statusCode)
    	}
    
    }
    ```
## 测试

- 传统的单元测试有如下缺点
    - 报错信息只有结果没有显示入参
    - 测试数据与测试逻辑混编在一起
    - 出现一个错误则测试流程结束
    ```java
    @Test
    public void testAdd() {
        assertEquals(3, add(1, 2));
        assertEquals(0, add(1, -1));
    }
    ```
- go 使用的是表格驱动测试
    - 能够自定义报错信息
    - 测试数据与测试逻辑分离
    - 一个错误不影响继续测试流程
- 单元测试的文件命名需要以 _test.go 结尾，如 add_test.go
- 命令行中使用 go test . 命令，运行当前目录下的所有测试
    ```go
    import (
    	"math"
    	"testing"
    )
    
    func TestAdd32(t *testing.T) {
    	tests := []struct{ a, b, c int32 }{
    		{1, 2, 3},
    		{0, 0, 0},
    		{math.MaxInt32, 1, math.MinInt32},
    	}
    	for _, test := range tests {
    		if actual := Add32(test.a, test.b); actual != test.c {
    			t.Errorf("add() error on a=%d and b=%d, expecting %d, actual got %d",
    				test.a, test.b, test.c, actual)
    		}
    	}
    }   
    ```
#### 测试覆盖率
    - 运行 go test 时，可以增加 -cover 参数，会打印覆盖率
    ```
    go test -cover .
    ```
    - 或者将覆盖率结果写到文件中，使用浏览器进行查看源码，使用不同颜色标记哪些覆盖了哪些没有覆盖
    ```
    go test -coverprofile=c.out .
    go tool cover -html=c.out
    ```
#### 性能测试
- 性能测试方法名以 Benchmark 开头，循环次数 b.N 由框架自己决定
- 执行如下命令可执行性能测试 go test -bench .
    ```go
    func Benchmark Add32(b *testing.B) {
    	var a1, a2 int32 = 1, 2
    	var res int32 = 3
    	for i := 0; i < b.N; i++ {
    		if actual := Add32(a1, a2); actual != res {
    			b.Errorf("add() error on a=%d and b=%d, expecting %d, actual got %d",
    				a1, a2, res, actual)
    		}
    	}
    }
    ```
- 查看性能消耗点
    - 到 http://graphviz.org/ 下载安装生成 svg 文件的工具
    ```
    // 生成分析结果
    go test -bench . -cpuprofile cpu.out
    
    // 进入交互式命令行工具
    go tool pprof cpu.out
    // 使用 web 方式，生成 svg 在浏览器查看
    (pprof) web
    ```
    - svg 图中，越大方框代表该逻辑点消耗的时间越多
    - 应该考虑多种输入的情况下，比较一个函数的写法（算法）的性能
    - string 转码为 rune 是需要消耗性能的
#### 测试 http
    ```go
    package webserver
    
    import (
    	"io/ioutil"
    	"net/http"
    	"net/http/httptest"
    	"strings"
    	"testing"
    )
    
    func createPanicErr(writer http.ResponseWriter, request *http.Request) error {
    	panic(123)
    }
    
    func createUserErr(writer http.ResponseWriter, request *http.Request) error {
    	return userErr("user error")
    }
    
    var tests = []struct {
    	h       requestHanlder
    	code    int
    	message string
    }{
    	{createPanicErr, 500, http.StatusText(http.StatusInternalServerError)},
    	{createUserErr, 400, "user error"},
    }
    
    // TestWebHandler 直接测试 handlerWrapper() 方法，使用假的 request 和 response
    func TestWebHandler(t *testing.T) {
    	for _, tt := range tests {
    		f := handlerWrapper(tt.h)
    		resp := httptest.NewRecorder()
    		req := httptest.NewRequest(
    			http.MethodGet,
    			"http://www.qq.com",
    			nil,
    		)
    		f(resp, req)
    
    		b, _ := ioutil.ReadAll(resp.Body)
    		body := strings.Trim(string(b), "\n")
    		if resp.Code != tt.code || body != tt.message {
    			t.Errorf("expect (%d, %s), got (%d, %s)", tt.code, tt.message,
    				resp.Code, body)
    		}
    	}
    }
    
    // TestWebServer 启动一个 web 服务器来测试
    func TestWebServer(t *testing.T) {
    	for _, tt := range tests {
    		f := handlerWrapper(tt.h)
    		server := httptest.NewServer(http.HandlerFunc(f))
    		resp, _ := http.Get(server.URL)
    
    		b, _ := ioutil.ReadAll(resp.Body)
    		body := strings.Trim(string(b), "\n")
    		if resp.StatusCode != tt.code || body != tt.message {
    			t.Errorf("expect (%d, %s), got (%d, %s)", tt.code, tt.message,
    				resp.StatusCode, body)
    		}
    	}
    }
    
    ```
## 文档工具

- go doc 命令查看具体包或结构体或方法的文档注释
- 安装 godoc，可在浏览器查看 go 官方文档和第三方用户代码文档
    - go env -w GO111MODULE=on
    - go get -v -u golang.org/x/tools/cmd/godoc
    - godoc -http :6060
- 示例代码
    - 实例代码的文件也需要用 _test.go 结尾
    - 实例方法以 Example 开头
    - 后面的 Output 注释带的是输出结果，运行 go test 时会进行校验
    - godoc 中也会将实例代码展示在 Example 区
    ```go
    package queue

    import "fmt"
    
    // ExampleIntQueue_Push 是对 IntQueue 的 Push() 方法的示例,
    // 测试时会检查注释中的 Output 是否正确,
    // godoc 中也会展示在 Example 区。
    func ExampleIntQueue_Push() {
    	q := IntQueue{1}
    	q.Push(12)
    	q.Push(22)
    	fmt.Println(q.IsEmpty())
    	fmt.Println(q.Pop())
    
    	// Output:
    	// false
    	// 1
    }
    ```
## goroutine 协程

- coroutine 协程是比子程序(subroutine)更广泛的感念，subroutine 是 coroutine 的特殊子集
    - 如下代码中，使用 go 关键字修饰了一个匿名函数的调用，这个匿名函数将使用协程进行运行。
    - 匿名函数的参数 index 是外层 for 循环的循环变量的副本，协程中如果不使用这个副本而直接使用 i 的话，将引发 race condition，由于闭包的关系，协程中使用的都是同一个变量 i。
    ```go
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
    ```
    - 协程是非抢占式多任务处理，由协程主动交出控制权
    - 是编译器/解释器/虚拟机层面的多任务，不是内核级别的多任务。（轻量级）
    - 多个协程可能在一个或多个线程上运行
- 使用如下命令运行程序，可检查是否由数据 race condition
    ```shell
    go run -race goroutineDemo.go
    ```
- main 方法也是在协程中执行的，如果 cpu 被其他协程占满不释放，那么 main 方法可能无法退出
- 其他语言的支持
    - C++ 通过 Boot.Coroutine 库支持
    - java 原生不支持，第三方 JVM 由支持
    - python 原来通过 yield 关键字进行支持，3.5 版本后通过 async def 进行支持。与 go 的区别是，python 在定义方法时就得标明是协程的，不能作为普通函数使用；而 go 定义时不需要做任何额外工作，只是在运行时使用 go 关键字修饰即可
    - JavaScript
- 一个线程中运行多少个协程，由调度器决定
- 协程运行的切换由调度器操作，可能切换的点：
    - I/O 操作，select
    - channel
    - 等待锁
    - 函数调用(有时)
    - runtime.Gosched()
- fmt.Printf() 是 I/O 操作，所以可以切换