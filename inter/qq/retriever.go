package qq

import (
	"net/http"
	"net/http/httputil"
	"time"
)

// Retriever qq 包中的 Retriever 实现类
type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

// Get 实现 Get 方法
func (r Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()
	return string(body)
}
