package errors

import "net/http"

type ErrCode struct {
	// 错误码
	C int

	// http的状态码
	HTTP int

	// 扩展字段
	Ext string

	// 引用文档
	Ref string
}

func (e ErrCode) Code() int {
	if e.C == 0 {
		return http.StatusInternalServerError
	}
	return e.C
}