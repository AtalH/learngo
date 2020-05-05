package main

import (
	"fmt"
	"math"
	"math/big"
	"math/cmplx"
)

func euler() {
	// 遵循浮点型规范的语言都会得到这个值，python 也一样
	// 0+1.2246467991473515e-16i
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)

	// 格式化之后能得到 0.000+0.000i
	fmt.Printf("euler result %.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

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

func main() {
	euler()
	bigNum()
}
