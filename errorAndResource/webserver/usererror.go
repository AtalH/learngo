package webserver

// UserError 自定义接口类型：用户错误
type UserError interface {
	error
	Message()
}
