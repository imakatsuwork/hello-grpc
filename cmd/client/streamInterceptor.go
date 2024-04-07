package main

import (
	"context"
	"errors"
	"io"
	"log"

	"google.golang.org/grpc"
)

// StreamClientInterceptor(https://github.com/grpc/grpc-go/blob/v1.63.0/interceptor.go#L63)を実装している
func customStreamClientInterceptor(
	ctx context.Context,
	desc *grpc.StreamDesc,
	cc *grpc.ClientConn,
	method string,
	streamer grpc.Streamer,
	opts ...grpc.CallOption,
) (grpc.ClientStream, error) {
	// ここではまだストリームが開始されていない
	log.Println("[ストリームの事前処理] custom stream client interceptor 1", method)

	stream, err := streamer(ctx, desc, cc, method, opts...)
	return &customClientStreamWrapper{stream}, err
}

// ClientでもStreamのメッセージ前後に処理を挟むために構造体をラップする

type customClientStreamWrapper struct {
	grpc.ClientStream
}

func (s *customClientStreamWrapper) SendMsg(m any) error {
	log.Println("[メッセージの処理前] custom stream client interceptor: ", m)

	// リクエスト送信
	return s.ClientStream.SendMsg(m)
}

func (s *customClientStreamWrapper) RecvMsg(m any) error {
	err := s.ClientStream.RecvMsg(m)

	// レスポンス受信後に割り込ませる処理
	if !errors.Is(err, io.EOF) {
		log.Println("[メッセージの処理後] custom stream client interceptor: ", m)
	}
	return err
}

func (s *customClientStreamWrapper) CloseSend() error {
	err := s.ClientStream.CloseSend()

	log.Println("[ストリームの事後処理] custom stream client interceptor")
	return err
}
