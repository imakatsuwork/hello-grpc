//　protoのバージョンの宣言
//　デフォルトは"proto2"

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: hello.proto

// packageの宣言(Goっぽい感じ)

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GreetingService_Hello_FullMethodName             = "/myapp.GreetingService/Hello"
	GreetingService_HelloServerStream_FullMethodName = "/myapp.GreetingService/HelloServerStream"
	GreetingService_HelloClientStream_FullMethodName = "/myapp.GreetingService/HelloClientStream"
	GreetingService_HelloBidiStream_FullMethodName   = "/myapp.GreetingService/HelloBidiStream"
	GreetingService_OccurError_FullMethodName        = "/myapp.GreetingService/OccurError"
)

// GreetingServiceClient is the client API for GreetingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetingServiceClient interface {
	// サービスが持つメソッドの定義(Procedureにあたるもの)
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	// サーバーストリーミング
	// ポイントは戻り値の型に「stream」がついていること
	HelloServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (GreetingService_HelloServerStreamClient, error)
	// クライアントストリーミング
	HelloClientStream(ctx context.Context, opts ...grpc.CallOption) (GreetingService_HelloClientStreamClient, error)
	// バイディレクショナルストリーミング
	HelloBidiStream(ctx context.Context, opts ...grpc.CallOption) (GreetingService_HelloBidiStreamClient, error)
	// 意図的にエラーを発生させる
	OccurError(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type greetingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetingServiceClient(cc grpc.ClientConnInterface) GreetingServiceClient {
	return &greetingServiceClient{cc}
}

func (c *greetingServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, GreetingService_Hello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetingServiceClient) HelloServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (GreetingService_HelloServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetingService_ServiceDesc.Streams[0], GreetingService_HelloServerStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &greetingServiceHelloServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetingService_HelloServerStreamClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greetingServiceHelloServerStreamClient struct {
	grpc.ClientStream
}

func (x *greetingServiceHelloServerStreamClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetingServiceClient) HelloClientStream(ctx context.Context, opts ...grpc.CallOption) (GreetingService_HelloClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetingService_ServiceDesc.Streams[1], GreetingService_HelloClientStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &greetingServiceHelloClientStreamClient{stream}
	return x, nil
}

type GreetingService_HelloClientStreamClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*HelloResponse, error)
	grpc.ClientStream
}

type greetingServiceHelloClientStreamClient struct {
	grpc.ClientStream
}

func (x *greetingServiceHelloClientStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetingServiceHelloClientStreamClient) CloseAndRecv() (*HelloResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetingServiceClient) HelloBidiStream(ctx context.Context, opts ...grpc.CallOption) (GreetingService_HelloBidiStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetingService_ServiceDesc.Streams[2], GreetingService_HelloBidiStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &greetingServiceHelloBidiStreamClient{stream}
	return x, nil
}

type GreetingService_HelloBidiStreamClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greetingServiceHelloBidiStreamClient struct {
	grpc.ClientStream
}

func (x *greetingServiceHelloBidiStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetingServiceHelloBidiStreamClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetingServiceClient) OccurError(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, GreetingService_OccurError_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetingServiceServer is the server API for GreetingService service.
// All implementations must embed UnimplementedGreetingServiceServer
// for forward compatibility
type GreetingServiceServer interface {
	// サービスが持つメソッドの定義(Procedureにあたるもの)
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	// サーバーストリーミング
	// ポイントは戻り値の型に「stream」がついていること
	HelloServerStream(*HelloRequest, GreetingService_HelloServerStreamServer) error
	// クライアントストリーミング
	HelloClientStream(GreetingService_HelloClientStreamServer) error
	// バイディレクショナルストリーミング
	HelloBidiStream(GreetingService_HelloBidiStreamServer) error
	// 意図的にエラーを発生させる
	OccurError(context.Context, *HelloRequest) (*HelloResponse, error)
	mustEmbedUnimplementedGreetingServiceServer()
}

// UnimplementedGreetingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetingServiceServer struct {
}

func (UnimplementedGreetingServiceServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedGreetingServiceServer) HelloServerStream(*HelloRequest, GreetingService_HelloServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloServerStream not implemented")
}
func (UnimplementedGreetingServiceServer) HelloClientStream(GreetingService_HelloClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloClientStream not implemented")
}
func (UnimplementedGreetingServiceServer) HelloBidiStream(GreetingService_HelloBidiStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloBidiStream not implemented")
}
func (UnimplementedGreetingServiceServer) OccurError(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OccurError not implemented")
}
func (UnimplementedGreetingServiceServer) mustEmbedUnimplementedGreetingServiceServer() {}

// UnsafeGreetingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetingServiceServer will
// result in compilation errors.
type UnsafeGreetingServiceServer interface {
	mustEmbedUnimplementedGreetingServiceServer()
}

func RegisterGreetingServiceServer(s grpc.ServiceRegistrar, srv GreetingServiceServer) {
	s.RegisterService(&GreetingService_ServiceDesc, srv)
}

func _GreetingService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetingServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GreetingService_Hello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetingServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetingService_HelloServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetingServiceServer).HelloServerStream(m, &greetingServiceHelloServerStreamServer{stream})
}

type GreetingService_HelloServerStreamServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type greetingServiceHelloServerStreamServer struct {
	grpc.ServerStream
}

func (x *greetingServiceHelloServerStreamServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetingService_HelloClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetingServiceServer).HelloClientStream(&greetingServiceHelloClientStreamServer{stream})
}

type GreetingService_HelloClientStreamServer interface {
	SendAndClose(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greetingServiceHelloClientStreamServer struct {
	grpc.ServerStream
}

func (x *greetingServiceHelloClientStreamServer) SendAndClose(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetingServiceHelloClientStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetingService_HelloBidiStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetingServiceServer).HelloBidiStream(&greetingServiceHelloBidiStreamServer{stream})
}

type GreetingService_HelloBidiStreamServer interface {
	Send(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greetingServiceHelloBidiStreamServer struct {
	grpc.ServerStream
}

func (x *greetingServiceHelloBidiStreamServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetingServiceHelloBidiStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetingService_OccurError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetingServiceServer).OccurError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GreetingService_OccurError_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetingServiceServer).OccurError(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GreetingService_ServiceDesc is the grpc.ServiceDesc for GreetingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "myapp.GreetingService",
	HandlerType: (*GreetingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _GreetingService_Hello_Handler,
		},
		{
			MethodName: "OccurError",
			Handler:    _GreetingService_OccurError_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HelloServerStream",
			Handler:       _GreetingService_HelloServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "HelloClientStream",
			Handler:       _GreetingService_HelloClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "HelloBidiStream",
			Handler:       _GreetingService_HelloBidiStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hello.proto",
}
