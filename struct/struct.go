package main

import "fmt"

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
