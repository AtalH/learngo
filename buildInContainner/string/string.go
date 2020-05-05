package main

import (
	"fmt"
	"unicode/utf8"
)

var chStr = "hi中国啊。"

var strBlock = `abc"qqq"
efd`

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

func main() {
	strContent()
	runeUtf8()
}
