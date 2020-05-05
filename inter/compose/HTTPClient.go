package compose

// Getter 定义 get 方法
type Getter interface {
	Get(url string) string
}

// Poster 定义 post 方法
type Poster interface {
	Post(url string) string
}

// HTTPClient 组合了 Get() Post() 方法，也可以定义自己的方法
type HTTPClient interface {
	Getter
	Poster
	Session(key string, value string) bool
}

// HTTPUtil 是 HTTPClient 接口的实现类
type HTTPUtil struct {
	url string
}

// Get 是 HTTPUtil 中的具体方法实现
func (http HTTPUtil) Get(url string) string {
	return "fake get result"
}

// Post 是 HTTPUtil 中的具体方法实现
func (http HTTPUtil) Post(url string) string {
	return "fake post result"
}

// Session 是 HTTPUtil 中的具体方法实现
func (http HTTPUtil) Session(key string, value string) bool {
	return true
}
