package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

// UnaryServerInterceptor(https://github.com/grpc/grpc-go/blob/f7c5e6a76259ec101fac0b1a6f62c05cc74c3d0c/interceptor.go#L87)を実装している
// Unaryって名前だけあって、Streamの時は実行されない
func customUnaryServerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	log.Println("[事前処理] custom unary interceptor: ", info.FullMethod)
	res, err := handler(ctx, req)
	log.Println("[事後処理] custom unary interceptor: ", res)
	return res, err
}
