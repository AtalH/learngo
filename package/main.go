package main

import (
	"fmt"

	"github.com/AtalH/learngo/package/queue"
	"github.com/AtalH/learngo/package/tree"
	"github.com/AtalH/learngo/package/xtree"
)

func main() {
	root := tree.Node{Value: 0}
	root.Print()

	xnode := xtree.Xnode{Node: &root}
	xnode.PostOrder()

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
