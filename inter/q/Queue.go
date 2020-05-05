package q

// Queue interface{}表示任何类型
type Queue []interface{}

// Push 入队
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

// Pop 出队
func (q *Queue) Pop() interface{} {
	tail := (*q)[0]
	*q = (*q)[1:]
	return tail
}

// IsEmpty 队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(*q) <= 0
}
