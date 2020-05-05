package tree

import "fmt"

// Node 树节点
/// 结构名，变量名，方法名首字母大写时，对其他包可见，即 public
type Node struct {
	Value       int
	Left, Right *Node
}

// Print 打印 value 值
func (node *Node) Print() {
	fmt.Println(node.Value)
}
