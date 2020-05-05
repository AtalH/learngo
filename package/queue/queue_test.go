package queue

import "fmt"

// ExampleIntQueue_Push 是对 IntQueue 的 Push() 方法的示例,
// 测试时会检查注释中的 Output 是否正确,
// godoc 中也会展示在 Example 区。
func ExampleIntQueue_Push() {
	q := IntQueue{1}
	q.Push(12)
	q.Push(22)
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())

	// Output:
	// false
	// 1
}
