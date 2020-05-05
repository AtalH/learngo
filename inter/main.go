package main

import (
	"fmt"

	"github.com/AtalH/learngo/inter/compose"
	"github.com/AtalH/learngo/inter/face"
	"github.com/AtalH/learngo/inter/qq"
	"github.com/AtalH/learngo/inter/sysface"
)

// getQQ 虽然参数是 face.Retriever，但是只要是有 Get 方法的都可以传进去
func getQQ(r face.Retriever) string {
	return r.Get("https://qq.com/")
}

func showInterface() {
	var retriever face.Retriever
	retriever = qq.Retriever{UserAgent: "Chrome", TimeOut: 10}
	fmt.Printf("type:%T, value:%v\n", retriever, retriever)
	// 当接口方法是指针接收者是，type就是指针类型，value就是地址类型
}

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

func typeAssertion() {
	var retriever face.Retriever
	retriever = qq.Retriever{UserAgent: "FireFox", TimeOut: 10}
	if r, ok := retriever.(qq.Retriever); ok {
		fmt.Println("this is a qq.Retriever, UserAgent:", r.UserAgent)
	} else {
		fmt.Printf("this is not a qq.Retriever, %T\n", r)
	}
}

func testInterfaceCompose() {
	var util compose.HTTPClient
	util = compose.HTTPUtil{}
	fmt.Println(util.Post("test"))
}

func testStringer() {
	s := sysface.Str{Content: "test 1"}
	fmt.Println("s to String:", s.String())
}

func main() {
	var retriever face.Retriever
	retriever = qq.Retriever{}
	page := getQQ(retriever)

	//简单的使用方法
	//page = getQQ(qq.Retriever{})
	fmt.Println(page)

	showInterface()
	typeSwith()
	typeAssertion()
	testInterfaceCompose()
	testStringer()
}
