// Code generated by Kitex v0.9.1. DO NOT EDIT.

package echo

import (
	"context"
	api "github.com/Blue-Berrys/GoMall/demo/demo_thrift/kitex_gen/api"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Echo(ctx context.Context, rep *api.Request, callOptions ...callopt.Option) (r *api.Response, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kEchoClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kEchoClient struct {
	*kClient
}

func (p *kEchoClient) Echo(ctx context.Context, rep *api.Request, callOptions ...callopt.Option) (r *api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Echo(ctx, rep)
}
