package face

// Retriever 定义了一个包含 Get 方法的接口
type Retriever interface {
	Get(url string) string
}
