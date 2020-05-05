// Package tree 树结构相关包
// 给 package 的注释要有固定开头
// 同一个目录下的所有文件的 package 都一样
package tree

// Traverse 中序遍历打印
// 方法的注释要以方法名开头
// public 的方法必须要有注释
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
