package sysface

import "fmt"

// Str 实现了 Stringer 接口
type Str struct {
	Content string
}

// String 实现 Stringer 接口的 String() 方法
func (s Str) String() string {
	return fmt.Sprintf("str's content is [%v]", s.Content)
}
