package traverse

import "fmt"

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
	root.onTraverse(func(n *Node) {
		//调用者可以自定义遍历时是输出节点值还是做其他操作
		fmt.Println(n.Value)
	})
}
