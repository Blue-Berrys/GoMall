package middlerware

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"time"
)

func Middleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resq interface{}) (err error) {
		begin := time.Now()
		err = next(ctx, req, resq)
		fmt.Println(time.Since(begin))
		return err
	}
}
