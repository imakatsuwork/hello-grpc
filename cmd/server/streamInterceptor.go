package main

import (
	"errors"
	"io"
	"log"

	"google.golang.org/grpc"
)

func customStreamServerInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("[ストリーム全体の事前処理] custom stream server interceptor: ", info.FullMethod)
	err := handler(srv, &customServerStreamWrapper{ss})
	log.Println("[ストリーム全体の事後処理] custom stream server interceptor: ")
	return err
}

// Streamのメッセージ前後に処理を挟むために構造体をラップする
// ServerStream(https://github.com/grpc/grpc-go/blob/f7c5e6a76259ec101fac0b1a6f62c05cc74c3d0c/stream.go#L1486)に基づいてオーバーライドしてる

type customServerStreamWrapper struct {
	grpc.ServerStream
}

func (s *customServerStreamWrapper) RecvMsg(m any) error {
	err := s.ServerStream.RecvMsg(m)
	if !errors.Is(err, io.EOF) {
		log.Println("[メッセージの処理前] custom stream server interceptor: ", m)
	}
	return err
}

func (s *customServerStreamWrapper) SendMsg(m any) error {
	log.Println("[メッセージの処理後] custom stream server interceptor: ", m)
	return s.ServerStream.SendMsg(m)
}
