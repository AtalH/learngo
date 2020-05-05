package xtree

import "github.com/AtalH/learngo/package/tree"

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
