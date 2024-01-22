package main

import (
	"context"
	"fmt"
	hellopb "grpc-sample/pkg/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	// 1. 8080番ポートでリッスンする
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// 2. gRPCサーバーを作成
	s := grpc.NewServer()

	// 3. gRPCサーバーにGreetingServiceを登録
	hellopb.RegisterGreetingServiceServer(s, NewMyServer())

	// 4. サーバーリフレクションの設定
	reflection.Register(s)

	// 5. 作成したgRPCサーバーを、8080番ポートで稼働させる
	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	// 6. Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()

}

func NewMyServer() hellopb.GreetingServiceServer {
	return &myServer{}
}

type myServer struct {
	// grpcが自動で生成してくれるまだ中身が実装されていない構造体
	// 未実装でもインターフェース満たすし、適当にエラーを返してくれてる
	hellopb.UnimplementedGreetingServiceServer
}

func (s *myServer) Hello(_ context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	return &hellopb.HelloResponse{
		Message: fmt.Sprintf("Hello %s!", req.Name),
	}, nil
}

func (s *myServer) HelloServerStream(
	req *hellopb.HelloRequest,
	stream hellopb.GreetingService_HelloServerStreamServer,
) error {
	resCount := 5
	for i := 0; i < resCount; i++ {
		// クライアントにレスポンスを送信(自動生成のコードがいい感じにラップしてくれてる)
		if err := stream.Send(&hellopb.HelloResponse{
			Message: fmt.Sprintf("[%d] Hello, %s!", i, req.GetName()),
		}); err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}

	// returnすればストリームも終了される(優秀)
	return nil
}
