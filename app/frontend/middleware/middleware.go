package middleware

import "github.com/cloudwego/hertz/pkg/app/server"

func Register(h *server.Hertz) { //运行的时候先注册
	h.Use(GlobalAuth())
}
