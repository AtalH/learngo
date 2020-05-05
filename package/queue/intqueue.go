package queue

// IntQueue 使用别名方式扩充slice实现整型队列
type IntQueue []int

// Push 入队
//     e.g. q.Push(123)
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
