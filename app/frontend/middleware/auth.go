package middleware

import (
	"context"
	frontendUtils "github.com/Blue-Berrys/GoMall/app/frontend/utlis"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 方便业务逻辑获取用户身份相关的
		// 从session中获取用户信息，放在context
		session := sessions.Default(c)
		if session != nil {
			var err error
			userId := session.Get("user_id")
			if err != nil || userId == nil {
				// 处理错误或nil值，例如记录日志或返回默认值
				c.Next(ctx)
				return
			}
			ctx = context.WithValue(ctx, frontendUtils.SessionUserId, userId) //返回一个新的 context，其中包含了 userId 的值
			c.Next(ctx)                                                       //将带有 userId 的新 context 传递下去，这样后续的处理中就可以从 context 中获取 userId
		} else {
			// 处理session为nil的情况
			c.Next(ctx)
		} //将带有 userId 的新 context 传递下去，这样后续的处理中就可以从 context 中获取 userId
	}
}

func Auth() app.HandlerFunc { //需要登录访问的页面用到的时候
	return func(ctx context.Context, c *app.RequestContext) {
		// 方便业务逻辑获取用户身份相关的
		// 从session中获取用户信息，放在context
		session := sessions.Default(c)
		userId := session.Get("user_id")
		if userId == nil {
			c.Redirect(302, []byte("/sign-in?next="+c.FullPath())) //c.FullPath() 是用于获取当前请求的路径
			c.Abort()                                              //这个方法会立即停止当前请求处理函数的执行。调用这个方法后，后续的处理器或中间件将不会再被执行。
			return
		}
		c.Next(ctx) //这个方法允许继续执行请求生命周期中的下一个中间件或处理器，如果不加Next就处理不了下一个中间件或处理器
	}
}
