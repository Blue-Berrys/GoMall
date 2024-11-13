package service

import (
	"context"
	api "github.com/Blue-Berrys/GoMall/demo/demo_thrift/kitex_gen/api"
	"testing"
)

func TestEcho_Run(t *testing.T) {
	ctx := context.Background()
	s := NewEchoService(ctx)
	// init req and assert value

	rep := &api.Request{}
	resp, err := s.Run(rep)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
