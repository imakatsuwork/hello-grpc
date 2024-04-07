package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

// UnaryClientInterceptor(https://github.com/grpc/grpc-go/blob/v1.63.0/interceptor.go#L43)を実装している
func customUnaryClientInterceptor(
	ctx context.Context,
	method string,
	req, res any,
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	fmt.Println("[事前処理] custom unary client interceptor", method, req)
	err := invoker(ctx, method, req, res, cc, opts...)
	fmt.Println("[事後処理] custom unary client interceptor", res)
	return err
}
