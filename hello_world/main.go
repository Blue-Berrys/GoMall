package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default()

	h.GET("/hello", func(ctx context.Context, c *app.RequestContext) {
		c.Data(consts.StatusOK, consts.MIMETextPlain, []byte("hello, world"))
	})

	h.Spin()
}
