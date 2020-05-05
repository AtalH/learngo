package webserver

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type userErr string

// Error 实现 error 的接口函数
func (e userErr) Error() string {
	return e.Message()
}

// Message 实现 UserError 接口的函数
func (e userErr) Message() string {
	return string(e)
}

const pathOfListFilePrefix string = "/list/"

// requestHanlder 定义一个拥有返回 error 的函数类型
type requestHanlder func(writer http.ResponseWriter, request *http.Request) error

// handlerWrapper 将 requestHanlder 包装，统一处理其错误信息，返回的函数符合 http.HandleFunc 要求
func handlerWrapper(hanler requestHanlder) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// recover 代码块，处理一些无法意料的错误
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		err := hanler(writer, request)
		if err == nil {
			return
		}
		log.Printf("handle request [%s] error: %s", request.URL, err.Error())
		// 对于可以展示给用户的 userErr，直接返回
		if ue, ok := err.(userErr); ok {
			http.Error(writer, ue.Message(), http.StatusBadRequest)
			return
		}
		statusCode := http.StatusOK
		switch {
		case os.IsNotExist(err):
			statusCode = http.StatusNotFound
		case os.IsPermission(err):
			statusCode = http.StatusForbidden
		default:
			statusCode = http.StatusInternalServerError
		}
		http.Error(writer, http.StatusText(statusCode), statusCode)
	}

}

func handleFileList(writer http.ResponseWriter, request *http.Request) error {
	// 截取 ulr 后的路径作为文件的路径
	url := request.URL.Path
	if strings.Index(url, pathOfListFilePrefix) != 0 {
		return userErr("path should start with " + pathOfListFilePrefix)
	}
	filepath := request.URL.Path[len(pathOfListFilePrefix):]
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(contents)
	return nil
}

// Start 启动服务器
func Start() {
	// 设定一个 url 及其处理函数
	http.HandleFunc("/", handlerWrapper(handleFileList))
	// 开启端口
	err := http.ListenAndServe(":8180", nil)
	if err != nil {
		panic(err)
	}
}
