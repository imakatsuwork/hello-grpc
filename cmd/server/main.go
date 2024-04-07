package main

import (
	"context"
	"errors"
	"fmt"
	hellopb "grpc-sample/pkg/grpc"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func main() {

	// 1. 8080番ポートでリッスンする
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// 2. gRPCサーバーを作成
	s := grpc.NewServer(
		// インターセプタが設定できる
		// grpcパッケージが優秀すぎる
		grpc.UnaryInterceptor(customUnaryServerInterceptor),

		// Note: この書き方はpanicを起こす
		//grpc.UnaryInterceptor(customUnaryServerInterceptor),
		//grpc.UnaryInterceptor(customUnaryServerInterceptor2),
		// 複数のUnaryInterceptorを設定する場合は、grpc.ChainUnaryInterceptorを使う
		//grpc.ChainUnaryInterceptor(customUnaryServerInterceptor, customUnaryServerInterceptor2),

		grpc.StreamInterceptor(customStreamServerInterceptor),
	)

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

func (s *myServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	// コンテキストからメタデータを取得
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Println("[メタデータ]", md)
	}

	// メタデータを設定
	headerMD := metadata.New(map[string]string{"type": "unary", "from": "server", "in": "header"})
	// grpc.SetHeaderはレスポンスの最初に送るヘッダーフレーム
	if err := grpc.SetHeader(ctx, headerMD); err != nil {
		return nil, err
	}

	trailerMD := metadata.New(map[string]string{"type": "unary", "from": "server", "in": "trailer"})
	// grpc.SetTrailerはレスポンスの最後に送るヘッダーフレーム
	if err := grpc.SetTrailer(ctx, trailerMD); err != nil {
		return nil, err
	}

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

func (s *myServer) HelloClientStream(stream hellopb.GreetingService_HelloClientStreamServer) error {
	nameList := make([]string, 0)
	for {
		// ストリームからリクエストを受け取っている
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			// ストリームの終端操作
			message := fmt.Sprintf("Hello, %v!", nameList)
			// クライアントにレスポンスを送信
			// サーバーストリームとは違って、明示的なクローズがされている
			return stream.SendAndClose(&hellopb.HelloResponse{
				Message: message,
			})
		}
		if err != nil {
			return err
		}
		nameList = append(nameList, req.GetName())
	}
}

func (s *myServer) HelloBidiStream(stream hellopb.GreetingService_HelloBidiStreamServer) error {
	// コンテキストからメタデータを取得
	// streamからコンテキストを取得する
	if md, ok := metadata.FromIncomingContext(stream.Context()); ok {
		log.Println("[メタデータ]", md)
	}

	// メタデータを設定
	// ヘッダーの送信タイミング:
	// 1. SetHeaderを使用するとすぐにヘッダーを送信できる
	// 2. 最初のメッセージが送信される時
	// 3. ステータスコードが送信される時
	headerMD := metadata.New(map[string]string{"type": "stream", "from": "server", "in": "header"})
	if err := stream.SetHeader(headerMD); err != nil {
		return err
	}

	// streamでSendHeaderを使用するとすぐにヘッダーを送信できる
	//if err := stream.SendHeader(headerMD); err != nil {
	//	return err
	//}

	// トレーラーはステータスコードが返却される時に送信される
	trailerMD := metadata.New(map[string]string{"type": "stream", "from": "server", "in": "trailer"})
	stream.SetTrailer(trailerMD)

	// returnすればストリームも終了

	for {
		// ストリームからリクエストを受け取っている
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}
		message := fmt.Sprintf("Hello, %v!", req.GetName())
		// クライアントにレスポンスを送信
		if err := stream.Send(&hellopb.HelloResponse{
			Message: message,
		}); err != nil {
			return err
		}
	}
}

func (s *myServer) OccurError(_ context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	stat := status.Newf(codes.Unknown, "unknown error occurred by %s", req.Name)
	stat, _ = stat.WithDetails(&errdetails.DebugInfo{
		Detail: "特に理由はないけどエラーです。",
	})
	err := stat.Err()
	return nil, err
}
