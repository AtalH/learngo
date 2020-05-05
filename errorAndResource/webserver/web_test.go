package webserver

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func createPanicErr(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

func createUserErr(writer http.ResponseWriter, request *http.Request) error {
	return userErr("user error")
}

var tests = []struct {
	h       requestHanlder
	code    int
	message string
}{
	{createPanicErr, 500, http.StatusText(http.StatusInternalServerError)},
	{createUserErr, 400, "user error"},
}

// TestWebHandler 直接测试 handlerWrapper() 方法
func TestWebHandler(t *testing.T) {
	for _, tt := range tests {
		f := handlerWrapper(tt.h)
		resp := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodGet,
			"http://www.qq.com",
			nil,
		)
		f(resp, req)

		b, _ := ioutil.ReadAll(resp.Body)
		body := strings.Trim(string(b), "\n")
		if resp.Code != tt.code || body != tt.message {
			t.Errorf("expect (%d, %s), got (%d, %s)", tt.code, tt.message,
				resp.Code, body)
		}
	}
}

// TestWebServer 启动一个 web 服务器来测试
func TestWebServer(t *testing.T) {
	for _, tt := range tests {
		f := handlerWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)

		b, _ := ioutil.ReadAll(resp.Body)
		body := strings.Trim(string(b), "\n")
		if resp.StatusCode != tt.code || body != tt.message {
			t.Errorf("expect (%d, %s), got (%d, %s)", tt.code, tt.message,
				resp.StatusCode, body)
		}
	}
}
