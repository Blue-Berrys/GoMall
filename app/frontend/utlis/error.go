package utlis

import "github.com/cloudwego/hertz/pkg/common/hlog"

func MustHandleError(err error) { // 处理错误
	if err != nil {
		hlog.Fatal(err)
	}
}
